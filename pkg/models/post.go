package models

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/dedpidgon/go-web-app/pkg/config"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	AuthorID   primitive.ObjectID `json:"_author" bson:"_author,omitempty"`
	AuthorName string             `json:"author" bson:"author,omitempty"`
	HasAuthor  bool               `json:"has_author" bson:"has_author"`
	Title      string             `json:"title" bson:"title,omitempty" validate:"required" min:"3" max:"200"`
	Content    string             `json:"content" bson:"content,omitempty" validate:"required" min:"5" max:"10000"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	Votes      []*Vote            `json:"-" bson:"votes"`
}

func (p *Post) UpdateFriend(f *Post) {
	if p.Title != "" {
		f.Title = p.Title
	}
	if p.Content != "" {
		f.Content = p.Content
	}

}

func (p *Post) exists() bool {
	filter := bson.M{
		"title":   p.Title,
		"content": p.Content,
	}
	if err := posts_collection.FindOne(ctx, filter).Decode(&p); err != nil {
		return false
	}
	return true
}

func (p *Post) Valid() []string {
	validation_errors := make([]string, 0)
	validate := validator.New()
	if p.exists() {
		validation_errors = append(validation_errors, "That post already exists")
	}
	err := validate.Struct(p)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			validation_errors = append(validation_errors, err.Error())
		}
	}
	return validation_errors
}

// crud
func (p *Post) Save() error {

	p.AuthorName = config.Client.Name
	p.AuthorID = config.Client.ID

	p.HasAuthor = true
	p.Votes = make([]*Vote, 0)
	res, err := posts_collection.InsertOne(ctx, p)
	if err != nil {
		return err
	}
	p.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func GetAllPosts() ([]*Post, error) {
	var posts []*Post
	cur, err := posts_collection.Find(ctx, bson.M{})
	if err != nil {
		return posts, err
	}
	err = cur.All(ctx, &posts)
	if err != nil {
		return posts, err
	}

	sort_posts(posts)
	return posts, nil
}
func (p *Post) Score() int32 {
	x := 0
	for _, vote := range p.Votes {
		x += int(vote.Value)
	}
	return int32(x)
}
func GetOnePost(id primitive.ObjectID) (*Post, error) {
	var post *Post
	err := posts_collection.FindOne(ctx, bson.M{"_id": id}).Decode(&post)
	return post, err
}

func (p *Post) Update() error {
	if p.AuthorID != config.Client.ID {
		return errors.New("you cant update someone else's post")
	}
	filter := bson.M{
		"_id": p.ID,
	}
	update := bson.M{
		"$set": p,
	}
	p.UpdatedAt = time.Now().UTC()
	_, err := posts_collection.UpdateOne(ctx, filter, update)
	return err
}
func (p *Post) Delete() error {
	if p.AuthorID != config.Client.ID {
		return errors.New("you cant delete someone else's post")
	}
	filter := bson.M{"_id": p.ID}
	_, err := posts_collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if _, err = votes_collection.DeleteMany(ctx, bson.M{"_parent": p.ID}); err != nil {
		return err
	}
	update := bson.M{"$set": bson.M{
		"has_parent": false,
		"_parent":    primitive.NilObjectID,
	}}
	filter = bson.M{"_parent": p.ID}
	_, err = comments_collection.UpdateMany(ctx, filter, update)
	return err

}

func (p *Post) Vote(v *Vote) error {
	if err := posts_collection.FindOne(ctx, bson.M{"_id": p.ID}).Decode(&p); err != nil {
		return err
	}
	// check if vote exists in your array
	for i, vote := range p.Votes {
		if vote.Author == v.Author {
			// user has already voted, changing value
			if vote.Value == v.Value {
				// user has erased their vote, remove from array
				x := p.Votes
				x[i] = x[len(x)-1]
				p.Votes = x[:len(x)-1]
			} else {
				vote.Value *= -1
			}
			return p.update_votes()
		}
	}
	// user is making a new vote
	v.CreatedAt = time.Now().UTC()
	v.UpdatedAt = time.Now().UTC()
	v.ID = primitive.NewObjectID()
	p.Votes = append(p.Votes, v)

	return p.update_votes()
}

func (p *Post) update_votes() error {
	update := bson.M{
		"$set": bson.M{
			"votes":      p.Votes,
			"updated_at": time.Now().UTC(),
		},
	}
	_, err := posts_collection.UpdateByID(ctx, p.ID, update)
	return err
}

func (p *Post) GetClient() *config.ClientData {
	return config.Client
}

func (p *Post) AsJsonString() string {
	b, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}

// template data
func (p *Post) Comments() []*Comment {
	var comments []*Comment
	cur, _ := comments_collection.Find(ctx, bson.M{"_parent": p.ID})
	cur.All(ctx, &comments)
	if comments != nil {
		sort_comments(comments)
		return comments
	} else {
		return []*Comment{}
	}
}

func (p *Post) UserHasVoted(id primitive.ObjectID) bool {
	for _, vote := range p.Votes {
		if vote.Author == id {
			return true
		}
	}
	return false
}
func (p *Post) IsUpvote(id primitive.ObjectID) bool {
	for _, vote := range p.Votes {
		if vote.Author == id {
			return vote.Value != -1
		}
	}

	return false
}

package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Author    primitive.ObjectID `json:"_author" bson:"_author,omitempty"`
	Title     string             `json:"title" bson:"title,omitempty" validate:"required" min:"3" max:"200"`
	Content   string             `json:"content" bson:"content,omitempty" validate:"required" min:"5" max:"10000"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	Errors    []string

	Score int32 `json:"score" bson:"-"`
}

// crud
func (p *Post) Save() error {
	validate := validator.New()
	err := validate.Struct(p)
	if err != nil {
		return err
	}
	res, err := posts_collection.InsertOne(ctx, p)
	if err != nil {
		return err
	}
	p.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func GetAllPosts() ([]*Post, error) {
	var posts []*Post
	cursor, err := posts_collection.Find(ctx, bson.M{})
	if err != nil {
		return posts, err
	}
	err = cursor.All(ctx, &posts)
	if err != nil {
		return posts, err
	}
	for _, post := range posts {
		calculate_score(post)
	}
	sort_posts(posts)
	return posts, nil
}
func GetOnePost(id primitive.ObjectID) (*Post, error) {
	var post *Post
	err := posts_collection.FindOne(ctx, bson.M{"_id": id}).Decode(&post)
	calculate_score(post)
	return post, err
}

func (p *Post) Update(x *Post) error {
	err := posts_collection.FindOneAndUpdate(ctx, bson.M{"_id": p.ID}, bson.M{"$set": x}).Decode(&p)
	return err
}
func (p *Post) Delete() error {
	_, err := posts_collection.DeleteOne(ctx, bson.M{"_id": p.ID})
	return err
}

func (p *Post) AsJsonString() string {
	b, err := bson.Marshal(p)
	if err != nil {
		return ""
	}
	return string(b)
}

// template data
func (p *Post) Comments() []*Comment {
	var comments []*Comment
	cur, err := comments_collection.Find(ctx, bson.M{"_parent": p.ID})
	if err != nil {
		p.Errors = append(p.Errors, err.Error())
	}
	if err = cur.All(ctx, &comments); err != nil {
		p.Errors = append(p.Errors, err.Error())
	}
	if comments != nil {
		sort_comments(comments)
		return comments
	} else {
		return []*Comment{}
	}
}

func (p *Post) add_to_score()              { p.Score += 1 }
func (p *Post) subtract_from_score()       { p.Score -= 1 }
func (p *Post) get_id() primitive.ObjectID { return p.ID }

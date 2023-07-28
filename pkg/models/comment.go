package models

import (
	"time"

	"github.com/dedpidgon/go-web-app/pkg/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	AuthorID   primitive.ObjectID `json:"_author" bson:"_author"`
	AuthorName string             `json:"author" bson:"author"`
	HasAuthor  bool               `json:"has_author" bson:"has_author"`
	Parent     primitive.ObjectID `json:"_parent" bson:"_parent"`
	HasParent  bool               `json:"has_parent" bson:"has_parent"`
	Content    string             `json:"content" bson:"content"`
	CreatedAt  time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt  time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	Score      int32              `json:"score" bson:"-"`
	Votes      []*Vote            `json:"_" bson:"votes"`
	OK         string             `json:"-" bson:"-"`
}

// crud
func (c *Comment) Save() error {

	c.AuthorID = config.Client.ID
	c.AuthorName = config.Client.Name
	c.HasParent = true
	c.HasAuthor = true

	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
	res, err := comments_collection.InsertOne(ctx, c)
	if err != nil {
		return err
	}

	c.ID = res.InsertedID.(primitive.ObjectID)
	return nil
}
func GetAllComments() ([]*Comment, error) {
	var comments []*Comment
	cursor, err := comments_collection.Find(ctx, bson.M{})
	if err != nil {
		return comments, err
	}
	err = cursor.All(ctx, &comments)
	if err != nil {
		return comments, err
	}
	for _, comment := range comments {
		comment.calculate_score()
	}
	return comments, nil
}
func GetOneComment(id primitive.ObjectID) (*Comment, error) {
	var comment *Comment
	err := comments_collection.FindOne(ctx, bson.M{"_id": id}).Decode(&comment)
	comment.calculate_score()
	return comment, err
}

func (c *Comment) Update(x *Post) error {
	err := comments_collection.FindOneAndUpdate(ctx, bson.M{"_id": c.ID}, bson.M{"$set": x}).Decode(&c)
	return err
}
func (c *Comment) Delete() error {
	_, err := posts_collection.DeleteOne(ctx, bson.M{"_id": c.ID})
	if err != nil {
		return err
	}
	filter := bson.M{"_parent": c.ID}
	_, err = votes_collection.DeleteMany(ctx, filter)
	return err
}

func (c *Comment) calculate_score() {
	for _, vote := range c.Votes {
		c.Score += int32(vote.Value)
	}
}

func (c *Comment) ModelType() string {
	return "comments"
}

func (c *Comment) GetClient() *config.ClientData {
	return config.Client
}

func (c *Comment) UserHasVoted(id primitive.ObjectID) bool {
	for _, vote := range c.Votes {
		if vote.Author == c.ID {
			return true
		}

	}
	return false
}
func (c *Comment) DateString() string {
	return c.CreatedAt.Format("3:04PM January 2 2006")
}
func (c *Comment) IsUpvote(id primitive.ObjectID) bool {
	for _, vote := range c.Votes {
		if vote.Author == c.ID {
			return vote.Value == 1
		}

	}
	return false
}

func (c *Comment) Vote(v *Vote) error {
	if err := comments_collection.FindOne(ctx, bson.M{"_id": c.ID}).Decode(&c); err != nil {
		return err
	}
	// check if vote exists in your array
	for i, vote := range c.Votes {
		if vote.Author == v.Author {
			// user has already voted, changing value
			c.OK = "update"
			if vote.Value == v.Value {
				// user has erased their vote, remove from array
				x := c.Votes
				x[i] = x[len(x)-1]
				c.Votes = x[:len(x)-1]
			} else {
				vote.Value *= -1
			}
			return c.update_votes()
		}
	}
	// user is making a new vote
	v.CreatedAt = time.Now().UTC()
	v.UpdatedAt = time.Now().UTC()
	v.ID = primitive.NewObjectID()
	c.Votes = append(c.Votes, v)
	if c.OK != "update" {
		c.OK = "bigfoot"
	}
	return c.update_votes()
}
func (c *Comment) update_votes() error {
	update := bson.M{
		"$set": bson.M{
			"votes":      c.Votes,
			"updated_at": time.Now(),
		},
	}
	_, err := posts_collection.UpdateByID(ctx, c.ID, update)
	return err
}

package models

import (
	"context"
	"sort"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var DatabaseName string = "thought_sea"
var ctx = context.TODO()
var validate *validator.Validate
var users_collection *mongo.Collection
var posts_collection *mongo.Collection
var comments_collection *mongo.Collection
var votes_collection *mongo.Collection
var Reader *TemplateReader

func Init(c *mongo.Client) {
	users_collection = c.Database(DatabaseName).Collection("users")
	posts_collection = c.Database(DatabaseName).Collection("posts")
	comments_collection = c.Database(DatabaseName).Collection("comments")
	votes_collection = c.Database(DatabaseName).Collection("votes")
	validate = validator.New()
	Reader = &TemplateReader{make([]string, 0)}
}

func sort_comments(comments []*Comment) {
	sort.Slice(comments, func(i, j int) bool {
		return comments[i].Score > comments[j].Score
	})
}
func sort_posts(posts []*Post) {
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Score() > posts[j].Score()
	})
}

type TemplateReader struct {
	Errors []string
}

func (g *TemplateReader) Posts() []*Post {
	var posts []*Post
	cur, err := posts_collection.Find(ctx, bson.M{})
	if err != nil {
		g.Errors = append(g.Errors, err.Error())
	}
	if err := cur.All(ctx, &posts); err != nil {
		g.Errors = append(g.Errors, err.Error())
	}
	sort_posts(posts)
	return posts
}
func (g *TemplateReader) Users() []*User {
	var users []*User
	cur, err := users_collection.Find(ctx, bson.M{})
	if err != nil {
		g.Errors = append(g.Errors, err.Error())
	}
	if err := cur.All(ctx, &users); err != nil {
		g.Errors = append(g.Errors, err.Error())
	}
	return users
}

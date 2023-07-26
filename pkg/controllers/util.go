package controllers

import (
	"github.com/dedpidgon/go-web-app/pkg/response"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Handle(r *gin.Engine) {
	v := &vote_controller{}
	u := &user_controller{}
	p := &post_controller{}
	c := &comment_controller{}

	v.use(r.Group("/api/votes"))
	u.use(r.Group("/api/users"))
	p.use(r.Group("/api/posts"))
	c.use(r.Group("/api/comments"))
}

func save_session(c *gin.Context, id primitive.ObjectID, name string) error {
	ses := sessions.Default(c)
	ses.Set("mongo_id", id)
	ses.Set("username", name)
	err := ses.Save()
	return err
}
func delete_session(c *gin.Context) error {
	ses := sessions.Default(c)
	ses.Clear()
	if err := ses.Save(); err != nil {
		response.ServerErr(c, err)
		return err
	}
	return nil
}

// func isFieldNil(field reflect.Value) bool {
// 	// Check if the field is the zero value for its type
// 	return reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface())
// }
// func merge_posts(x, y *models.Post) {
// 	post := reflect.ValueOf(x)
// 	update := reflect.ValueOf(y)
// 	for i := 0; i < update.NumField(); i++ {
// 		field := update.Field(i)
// 		fieldName := update.Type().Field(i).Name

// 		// Check if the field is nil or not
// 		if isFieldNil(field) {

// 		} else {
// 			post.FieldByName(fieldName) = update.Field(i)
// 		}
// 	}
// }

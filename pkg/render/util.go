package render

import (
	"bytes"
	"fmt"
	"html/template"
	"path/filepath"

	"github.com/dedpidgon/go-web-app/pkg/controllers"
	"github.com/dedpidgon/go-web-app/pkg/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const path = "./templates/"

var cache = make(map[string]*template.Template)
var data *stuff

func standard_stuff() *stuff {
	return &stuff{
		controllers.UserController,
		controllers.PostController,
		nil,
		nil,
		"",
		limit_text,
	}
}

type stuff struct {
	user_view
	post_view
	U         *models.User
	P         *models.Post
	X         string
	LimitText func(string, int) string
}

type user_view interface {
	UserErrors() []string
	IsAuthenticated() bool
	UserName() string
	UserID() primitive.ObjectID
	GetAllUsers() []*models.User
	GetIP() string
}
type post_view interface {
	GetAllPosts() []*models.Post
	PostErrors() []string
}

type template_error struct {
	text string
}

func (te *template_error) Error() string {
	return te.text
}
func new_error(t string) *template_error {
	return &template_error{t}
}

func render_template(c *gin.Context, t string) error {
	name := fmt.Sprintf("%s.page.html", t)

	_, ok := cache[name]
	if !ok {
		return new_error("no template to render")
	}

	buf := new(bytes.Buffer)
	if err := cache[name].Execute(buf, data); err != nil {
		return err
	}

	if _, err := buf.WriteTo(c.Writer); err != nil {
		return err
	}
	return nil
}

func create_template_cache() error {

	pages, err := filepath.Glob(path + "*.page.html")
	if err != nil {
		return err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return err
		}

		layouts, err := filepath.Glob(path + "*.layout.html")
		if err != nil {
			return err
		}
		if len(layouts) > 0 {
			ts, err = ts.ParseGlob(path + "*.layout.html")
			if err != nil {
				return err
			}
		}

		partials, err := filepath.Glob(path + "*.partial.html")
		if err != nil {
			return err
		}

		if len(partials) > 0 {
			ts, err = ts.ParseGlob(path + "*.partial.html")
			if err != nil {
				return err
			}
		}
		cache[name] = ts

	}
	return nil
}

func limit_text(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}
	return s[:maxLength] + "..."
}

package render

import (
	"bytes"
	"fmt"
	"github.com/dedpidgon/go-web-app/pkg/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"html/template"
	"io"
	"path/filepath"
)

type template_data struct {
	user_view
	post_view
	U *models.User
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

func render_template(w io.Writer, t string, data interface{}) error {
	name := fmt.Sprintf("%s.page.gohtml", t)

	_, ok := cache[name]
	if !ok {
		return new_error("no template to render")
	}

	buf := new(bytes.Buffer)
	if err := cache[name].Execute(buf, data); err != nil {
		return err
	}

	if _, err := buf.WriteTo(w); err != nil {
		return err
	}
	return nil
}

func create_template_cache() error {

	pages, err := filepath.Glob(path + "*.page.gohtml")
	if err != nil {
		return err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return err
		}

		layouts, err := filepath.Glob(path + "*.layout.gohtml")
		if err != nil {
			return err
		}
		if len(layouts) > 0 {
			ts, err = ts.ParseGlob(path + "*.layout.gohtml")
			if err != nil {
				return err
			}
		}

		partials, err := filepath.Glob(path + "*.partial.gohtml")
		if err != nil {
			return err
		}

		if len(partials) > 0 {
			ts, err = ts.ParseGlob(path + "*.partial.gohtml")
			if err != nil {
				return err
			}
		}
		cache[name] = ts

	}
	return nil
}

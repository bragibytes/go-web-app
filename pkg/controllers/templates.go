package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/dedpidgon/go-web-app/pkg/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"html/template"
	"io"
	"log"
	"path/filepath"
)

type tmpl_user_controller interface {
	GetAll() []*models.User
	Authenticated() bool
	UserName() string
	UserID() primitive.ObjectID
}

type template_controller struct {
	path  string
	cache map[string]*template.Template
	data  *template_data
}

type template_data struct {
	IP string
	*user_controller
}

func new_template_data() *template_data {

	x := &template_data{
		"",
		UserController,
	}
	return x
}

func (tc *template_controller) use(r *gin.Engine) {
	r.GET("/", tc.home_page)
}

func new_template_controller() *template_controller {
	x := &template_controller{}
	x.path = "./templates/"
	x.cache = make(map[string]*template.Template)
	x.data = new_template_data()
	if err := x.create_template_cache(); err != nil {
		log.Fatal(err.Error())
	}
	return x
}

func (tc *template_controller) home_page(c *gin.Context) {

	if err := tc.render_template(c.Writer, "home", UserController); err != nil {
		response{
			false,
			"could not render the template",
			err.Error(),
			nil,
			500}.send(c)
		return
	}

}

func (tc *template_controller) render_template(w io.Writer, t string, data interface{}) error {
	name := fmt.Sprintf("%s.page.gohtml", t)

	_, ok := tc.cache[name]
	if !ok {
		return errors.New("no template to render")
	}

	buf := new(bytes.Buffer)
	if err := tc.cache[name].Execute(buf, data); err != nil {
		return err
	}

	if _, err := buf.WriteTo(w); err != nil {
		return err
	}
	return nil
}

func (tc *template_controller) create_template_cache() error {

	pages, err := filepath.Glob(tc.path + "*.page.gohtml")
	if err != nil {
		log.Println("--------!!!! ERROR !!!!-------------", err.Error(), " line 55")
		return err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			log.Println("--------!!!! ERROR !!!!-------------", err.Error(), " line 63")
			return err
		}

		layouts, err := filepath.Glob(tc.path + "*.layout.gohtml")
		if err != nil {
			log.Println("--------!!!! ERROR !!!!-------------", err.Error(), " line 69")
			return err
		}
		if len(layouts) > 0 {
			ts, err = ts.ParseGlob(tc.path + "*.layout.gohtml")
			if err != nil {
				log.Println("--------!!!! ERROR !!!!-------------", err.Error(), " line 75")
				return err
			}
		}

		partials, err := filepath.Glob(tc.path + "*.partial.gohtml")
		if err != nil {
			log.Println("--------!!!! ERROR !!!!-------------", err.Error(), " line 82")
			return err
		}

		if len(partials) > 0 {
			ts, err = ts.ParseGlob(tc.path + "*.partial.gohtml")
			if err != nil {
				log.Println("--------!!!! ERROR !!!!-------------", err.Error(), " line 89")
				return err
			}
		}
		tc.cache[name] = ts

	}
	return nil
}
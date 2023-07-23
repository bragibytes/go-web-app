package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
	"log"
	"path/filepath"
)

type template_controller struct {
	path  string
	cache map[string]*template.Template
}

func (tc *template_controller) use(r *gin.Engine) {
	r.GET("/", tc.home_page)
}

func new_template_controller() *template_controller {
	x := &template_controller{}
	x.path = "./templates/"
	x.cache = make(map[string]*template.Template)
	if err := x.create_template_cache(); err != nil {
		log.Fatal(err.Error())
	}
	return x
}

func (tc *template_controller) home_page(c *gin.Context) {

	if err := tc.render_template(c.Writer, "home", UserController); err != nil {
		response{
			"error",
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
		return err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return err
		}

		layouts, err := filepath.Glob(tc.path + "*.layout.gohtml")
		if err != nil {
			return err
		}
		if len(layouts) > 0 {
			ts, err = ts.ParseGlob(tc.path + "*.layout.gohtml")
			if err != nil {
				return err
			}
		}

		partials, err := filepath.Glob(tc.path + "*.partial.gohtml")
		if err != nil {
			return err
		}

		if len(partials) > 0 {
			ts, err = ts.ParseGlob(tc.path + "*.partial.gohtml")
			if err != nil {
				return err
			}
		}
		tc.cache[name] = ts

	}
	return nil
}

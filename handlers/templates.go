package handlers

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

const path string = "./templates/"

var templateCache = make(map[string]*template.Template)

func home_page(c *gin.Context) {
	if err := render_template(c.Writer, "home", nil); err != nil {
		response{
			false,
			"could not render the template",
			err.Error(),
			nil,
			500}.send(c)
		return
	}

}

func render_template(w io.Writer, t string, data interface{}) error {
	name := fmt.Sprintf("%s.page.gohtml", t)

	_, ok := templateCache[name]
	if !ok {
		return errors.New("no template to render")
	}

	buf := new(bytes.Buffer)
	if err := templateCache[name].Execute(buf, data); err != nil {
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

		layouts, err := filepath.Glob(path + "*.layout.gohtml")
		if err != nil {
			log.Println("--------!!!! ERROR !!!!-------------", err.Error(), " line 69")
			return err
		}
		if len(layouts) > 0 {
			ts, err = ts.ParseGlob(path + "*.layout.gohtml")
			if err != nil {
				log.Println("--------!!!! ERROR !!!!-------------", err.Error(), " line 75")
				return err
			}
		}

		partials, err := filepath.Glob(path + "*.partial.gohtml")
		if err != nil {
			log.Println("--------!!!! ERROR !!!!-------------", err.Error(), " line 82")
			return err
		}

		if len(partials) > 0 {
			ts, err = ts.ParseGlob(path + "*.partial.gohtml")
			if err != nil {
				log.Println("--------!!!! ERROR !!!!-------------", err.Error(), " line 89")
				return err
			}
		}
		templateCache[name] = ts

	}
	return nil
}

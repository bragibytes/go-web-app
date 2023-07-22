package views

import (
	"backend/controllers"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"io"
)

const path string = "./views/"

var users *controllers.UserController

var templateCache = make(map[string]*template.Template)

func Init(r *gin.Engine) {
	users = controllers.NewUserController()
	// Pages
	r.GET("/", users.Middleware(), home)

}
func home(c *gin.Context) {
	if err := renderTemplate(c.Writer, "home", users); err != nil {
		c.JSON(500, err.Error())
	}
}

func renderTemplate(w io.Writer, t string, data interface{}) error {
	var tmpl *template.Template
	var err error

	name := fmt.Sprintf(path+"%s.page.gohtml", t)
	if _, inCache := templateCache[name]; !inCache {
		err = addToCache(t)
	}

	tmpl = templateCache[t]
	err = tmpl.Execute(w, data)
	return err
}

func addToCache(t string) error {
	templates := []string{
		fmt.Sprintf(path+"%s.page.gohtml", t),
		layout("base"),
		component("header"),
		component("user"),
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	templateCache[t] = tmpl
	return nil
}

func component(n string) string {
	return fmt.Sprintf(path+"%s.component.gohtml", n)
}
func layout(n string) string {
	return fmt.Sprintf(path+"%s.layout.gohtml", n)
}

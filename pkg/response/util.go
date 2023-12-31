package response

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	success  = "success"
	warning  = "warning"
	breaking = "error"
	info     = "info"
)

type response struct {
	MessageType string      `json:"message_type"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func Send(c *gin.Context, t string, m string, d interface{}, n int) {
	x := &response{t, m, d}
	c.JSON(n, x)
}

func ValidationErrors(c *gin.Context, errors []string) {
	res := &response{
		warning,
		"Validation Errors",
		errors,
	}
	c.JSON(http.StatusTeapot, res)
}

func BadReq(c *gin.Context, err error) {
	res := &response{
		warning,
		err.Error(),
		nil,
	}
	c.JSON(http.StatusBadRequest, res)
}
func ServerErr(c *gin.Context, err error) {
	res := &response{
		breaking,
		err.Error(),
		nil,
	}
	c.JSON(http.StatusInternalServerError, res)
}
func Unauthorized(c *gin.Context, msg string) {
	res := &response{
		warning,
		capitalize_words(msg),
		nil,
	}
	c.JSON(http.StatusUnauthorized, res)
}
func OK(c *gin.Context, msg string, data interface{}) {
	res := &response{
		success,
		capitalize_words(msg),
		data,
	}
	c.JSON(http.StatusOK, res)
}
func Created(c *gin.Context, m string, data interface{}) {
	res := &response{
		success,
		capitalize(m),
		data,
	}
	c.JSON(http.StatusCreated, res)
}
func NotFound(c *gin.Context, x string, err string) {
	res := &response{
		warning,
		capitalize_words("could not find the " + x + " you were looking for."),
		err,
	}
	c.JSON(http.StatusNotFound, res)
}

func capitalize(s string) string {
	return strings.ToUpper(s[:1]) + s[1:]
}

func capitalize_words(s string) string {
	words := strings.Fields(s)
	for i, word := range words {
		words[i] = capitalize(word)
	}
	return strings.Join(words, " ")
}

package response

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type response struct {
	MessageType string      `json:"message_type"` //'success', 'warning', 'error', 'info', 'neutral']
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

func Send(c *gin.Context, t string, m string, d interface{}, n int) {
	x := &response{t, m, d}
	c.JSON(n, x)
}

func BadReq(c *gin.Context, err error) {
	res := &response{
		"warning",
		err.Error(),
		nil,
	}
	c.JSON(http.StatusBadRequest, res)
}
func ServerErr(c *gin.Context, err error) {
	res := &response{
		"error",
		err.Error(),
		nil,
	}
	c.JSON(http.StatusInternalServerError, res)
}
func Unauthorized(c *gin.Context, msg string) {
	res := &response{
		"warning",
		capitalize_words(msg),
		nil,
	}
	c.JSON(http.StatusUnauthorized, res)
}
func OK(c *gin.Context, msg string, data interface{}) {
	res := &response{
		"success",
		capitalize_words(msg),
		data,
	}
	c.JSON(http.StatusOK, res)
}
func Created(c *gin.Context, t string, data interface{}) {
	res := &response{
		"success",
		capitalize(t) + " Created!",
		data,
	}
	c.JSON(http.StatusCreated, res)
}
func NotFound(c *gin.Context, x string, err string) {
	res := &response{
		"error",
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

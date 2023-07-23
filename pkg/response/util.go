package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type response struct {
	MessageType string      `json:"message_type"` //'success', 'warning', 'error', 'info', 'neutral']
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
	Code        int         `json:"code"`
}

func BadReq(c *gin.Context, err error) {
	res := &response{
		"warning",
		err.Error(),
		nil,
		http.StatusBadRequest,
	}
	c.JSON(res.Code, res)
}
func ServerErr(c *gin.Context, err error) {
	res := &response{
		"error",
		err.Error(),
		nil,
		http.StatusInternalServerError,
	}
	c.JSON(res.Code, res)
}
func Unauthorized(c *gin.Context, msg string) {
	res := &response{
		"warning",
		capitalize_words(msg),
		nil,
		http.StatusUnauthorized,
	}
	c.JSON(res.Code, res)
}
func OK(c *gin.Context, msg string, data interface{}) {
	res := &response{
		"success",
		capitalize_words(msg),
		data,
		http.StatusOK,
	}
	c.JSON(res.Code, res)
}
func Created(c *gin.Context, t string, data interface{}) {
	res := &response{
		"success",
		capitalize(t) + " Created!",
		data,
		http.StatusCreated,
	}
	c.JSON(res.Code, res)
}
func NotFound(c *gin.Context, x string, err string) {
	res := &response{
		"error",
		capitalize_words("could not find the " + x + " you were looking for."),
		err,
		http.StatusNotFound,
	}
	c.JSON(res.Code, res)
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

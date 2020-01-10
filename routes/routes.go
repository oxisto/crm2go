package routes

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO: move to httputil
type ErrorResponse struct {
	Error string `json:"error"`
}

func JSON(c *gin.Context, status int, value interface{}, err error) {
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{err.Error()})
		return
	}

	if value == nil || (reflect.ValueOf(value).Kind() == reflect.Ptr && reflect.ValueOf(value).IsNil()) {
		c.JSON(http.StatusNotFound, nil)
	} else {
		c.JSON(status, value)
	}
}

func IntParam(c *gin.Context, key string) (i int64, err error) {
	return strconv.ParseInt(c.Param(key), 10, 64)
}

// NewRouter returns a configured router containing all REST endpoints
func NewRouter() *gin.Engine {
	r := gin.Default()

	api := r.Group("/api")
	/*api.Use(handler.AuthRequired)
	{*/
	contacts := api.Group("/contacts")
	/*persons.Use(AccountRequired)
	{*/
	contacts.GET("/", GetAllContacts)
	contacts.GET("/:id", GetContact)

	// @todo Replace the post requests with a message queue

	contacts.POST("/", PostContact)
	contacts.PUT("/:id", PutContact)
	//}
	//}

	return r
}

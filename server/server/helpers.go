package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mayudev/notesplace/server/util"
)

// bindData binds data
func bindData[T any](c *gin.Context, data T) bool {
	if err := c.ShouldBindJSON(&data); err != nil {
		return false
	}

	return true
}

// internalServerError returns a 500 Internal Server Error
func internalServerError(c *gin.Context) {
	c.JSON(500, util.Response{
		Status:  "error",
		Message: util.InternalServerError,
	})
}

// invalidRequest returns a 400 Bad Request error
func badRequest(c *gin.Context) {
	c.JSON(400, util.Response{
		Status:  "error",
		Message: util.RequestInvalid,
	})
}

// notFound returns a 404 Not Found error
func notFound(c *gin.Context) {
	c.JSON(404, util.Response{
		Status:  "error",
		Message: util.NotFound,
	})
}

// unauthorized returns a 401 Unauthorized error
func unauthorized(c *gin.Context) {
	c.JSON(401, util.Response{
		Status:  "error",
		Message: util.Unauthorized,
	})
}

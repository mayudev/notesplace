package server

import (
	"strings"

	"github.com/gin-gonic/gin"
)

// Validate validates if token provided in the request's header
// is valid for provided notebook ID
func (s *Server) Validate(c *gin.Context, notebook string) bool {
	authorization := c.GetHeader("Authorization")
	if len(authorization) == 0 {
		unauthorized(c)
		return false
	}

	token := strings.TrimPrefix(authorization, "Bearer ")
	valid := s.issuer.ValidateNotebook(token, notebook)

	return valid
}

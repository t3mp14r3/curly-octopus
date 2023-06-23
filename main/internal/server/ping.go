package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) ping(c *gin.Context) {
    login, _ := c.Get("login")

    c.JSON(http.StatusOK, gin.H{"login": login})
}

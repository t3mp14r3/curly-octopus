package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) me(c *gin.Context) {
    userIDAny, _ := c.Get("userID")
    userID := fmt.Sprint(userIDAny)

    user, err := s.repo.GetUser(s.ctx, userID)

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    user.Password = ""

    c.JSON(http.StatusOK, user)
}

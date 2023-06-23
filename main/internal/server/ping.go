package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) ping(c *gin.Context) {
    loginAny, _ := c.Get("login")
    login := fmt.Sprint(loginAny)

    user, err := s.repo.GetUserByLogin(s.ctx, login)

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    user.Password = ""

    c.JSON(http.StatusOK, user)
}

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) Ping(c *gin.Context) {
    token := c.GetHeader("Authorization")

    if len(token) == 0 {
        s.logger.Error("failed to get authorization token from header")
        c.JSON(http.StatusBadRequest, gin.H{"error": "authorization token is required"})
        return
    }

    ok := s.auth.Validate(s.ctx, token)
    
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
        return
    }

    login, err := s.auth.Extract(s.ctx, token)

    if err != nil  {
        c.JSON(http.StatusBadRequest, gin.H{"error": "failed to extract data from token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"login": login})
}

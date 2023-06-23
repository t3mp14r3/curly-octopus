package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) withAuth(c *gin.Context) {
    token := c.GetHeader("Authorization")

    if len(token) == 0 {
        s.logger.Error("failed to get authorization token from header")
        c.JSON(http.StatusBadRequest, gin.H{"error": "authorization token is required"})
        c.Abort()
        return
    }

    ok := s.auth.Validate(s.ctx, token)
    
    if !ok {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
        c.Abort()
        return
    }

    userID, err := s.auth.Extract(s.ctx, token)

    if err != nil  {
        c.JSON(http.StatusBadRequest, gin.H{"error": "failed to extract data from token"})
        c.Abort()
        return
    }

    c.Set("userID", userID)

    c.Next()
}

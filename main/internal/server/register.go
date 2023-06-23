package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *Server) Register(c *gin.Context) {
    body := make(map[string]string)
    err := c.BindJSON(&body)

    if err != nil {
        s.logger.Error("failed to parse request body", zap.Error(err))
        c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse request body"})
        return
    }

    if _, ok := body["login"]; !ok {
        s.logger.Error("login field not found", zap.Error(err))
        c.JSON(http.StatusBadRequest, gin.H{"error": "login frield not found"})
        return
    }

    token, err := s.auth.Generate(s.ctx, body["login"])
    
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate the token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

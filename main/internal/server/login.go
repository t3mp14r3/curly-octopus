package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/t3mp14r3/curly-octopus/main/internal/domain"
	"go.uber.org/zap"
    "golang.org/x/crypto/bcrypt"
)

func (s *Server) login(c *gin.Context) {
    var input domain.LoginRequest
    err := c.BindJSON(&input)

    if err != nil {
        s.logger.Error("failed to parse request body", zap.Error(err))
        c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse request body"})
        return
    }

    if err := input.Validate(); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := s.repo.GetUserByLogin(s.ctx, input.Login)
    
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

    if err != nil {
        s.logger.Error("failed to verify password hash", zap.Error(err))
        c.JSON(http.StatusForbidden, gin.H{"error": "incorrect password"})
        return
    }

    token, err := s.auth.Generate(s.ctx, input.Login)
    
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate the token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/t3mp14r3/curly-octopus/main/internal/domain"
	"go.uber.org/zap"
    "golang.org/x/crypto/bcrypt"
)

func (s *Server) register(c *gin.Context) {
    var input domain.RegisterRequest
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

    if s.repo.EmailUsed(s.ctx, input.Email) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "email is already in use"})
        return
    }
    
    if s.repo.LoginUsed(s.ctx, input.Email) {
        c.JSON(http.StatusBadRequest, gin.H{"error": "login is already in use"})
        return
    }

    bytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

    if err != nil {
        s.logger.Error("failed to generate password hash", zap.Error(err))
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate password hash"})
        return
    }

    newUser := domain.User{
        Login: input.Login,
        Name: input.Name,
        Email: input.Email,
        Password: string(bytes),
    }

    user, err := s.repo.CreateUser(s.ctx, newUser)
    
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create new user record"})
        return
    }

    token, err := s.auth.Generate(s.ctx, user.ID)
    
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate the token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

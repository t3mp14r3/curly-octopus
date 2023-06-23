package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/t3mp14r3/curly-octopus/main/internal/domain"
	"go.uber.org/zap"
)

func (s *Server) createProduct(c *gin.Context) {
    userIDAny, _ := c.Get("userID")
    userID := fmt.Sprint(userIDAny)

    var input domain.CreateProductRequest
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

    newProduct := domain.Product{
        Name: input.Name,
        Desc: input.Desc,
        Cost: input.Cost,
        Barcode: input.Barcode,
        UserID: userID,
    }

    product, err := s.repo.CreateProduct(s.ctx, newProduct)
    
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create new product"})
        return
    }

    c.JSON(http.StatusOK, product)
}

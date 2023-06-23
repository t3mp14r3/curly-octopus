package server

import (
	"fmt"
	"net/http"

    "github.com/google/uuid"
	"github.com/gin-gonic/gin"
)

func (s *Server) getProduct(c *gin.Context) {
    userIDAny, _ := c.Get("userID")
    userID := fmt.Sprint(userIDAny)

    productID := c.Param("id")
    
    if len(productID) == 0 {
        c.JSON(http.StatusBadRequest, gin.H{"error": "product id is requried"})
        return
    }

    _, err := uuid.Parse(productID)

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
        return
    }

    product, err := s.repo.GetProduct(s.ctx, productID)

    if (product == nil && err == nil) || product.UserID != userID {
        c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
        return
    }
    
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get product record"})
        return
    }

    c.JSON(http.StatusOK, product)
}

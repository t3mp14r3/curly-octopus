package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) getProducts(c *gin.Context) {
    userIDAny, _ := c.Get("userID")
    userID := fmt.Sprint(userIDAny)

    products, err := s.repo.GetUserProducts(s.ctx, userID)

    if products == nil && err == nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "no products found"})
        return
    }
    
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user products"})
        return
    }

    c.JSON(http.StatusOK, products)
}

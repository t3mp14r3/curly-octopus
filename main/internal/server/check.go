package server

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/t3mp14r3/curly-octopus/checks/gen"
)

func (s *Server) check(c *gin.Context) {
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

    req := gen.CreateRequest{
        Id: product.ID,
        Name: product.Name,
        Barcode: product.Barcode,
        Cost: product.Cost,
    }

    resp, err := s.checks.Create(s.ctx, &req)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate the check"})
        return
    }

    c.Header("Content-Type", "application/pdf")
    http.ServeContent(c.Writer, c.Request, resp.Filename, time.Now(), bytes.NewReader(resp.Data))
}

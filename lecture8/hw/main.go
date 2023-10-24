package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

const AuthToken = "tokenXcxzcasdKLDSAdxc"

type Item struct {
	ID    int    `json:"id"`
	Name  string `json:"name" binding:"required"`
	Price int    `json:"price"`
}

var items = []Item{
	{ID: 1, Name: "Item 1", Price: 10},
	{ID: 2, Name: "Item 2", Price: 20},
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		fmt.Println("Received token:", token)
		if token != AuthToken {
			c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func main() {
	r := gin.Default()

	r.Use(AuthMiddleware())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/items", func(c *gin.Context) {
			c.JSON(http.StatusOK, items)
		})

		v1.POST("/items", func(c *gin.Context) {
			var newItem Item
			if err := c.ShouldBindWith(&newItem, binding.JSON); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			items = append(items, newItem)
			c.JSON(http.StatusCreated, newItem)
		})
	}

	r.Run(":3000")
}

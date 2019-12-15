package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addOrder(c *gin.Context) {
	var order Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(order.Products) < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Order has no products!"})
		return
	}

	store.addOrder(&order)

	c.JSON(http.StatusOK, gin.H{
		"orderId": order.ID,
	})
}

func getOrder(c *gin.Context) {
	id := c.Param("id")

	if order, ok := store.getOrder(id); ok {
		c.JSON(http.StatusOK, gin.H{
			"status": order.Status.String(),
		})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"orderId": id,
	})
}

func deleteOrder(c *gin.Context) {
	id := c.Param("id")

	if _, ok := store.deleteOrder(id); ok {
		c.Status(http.StatusOK)
		return
	}

	c.JSON(http.StatusNotFound, gin.H{
		"orderId": id,
	})
}

func setupRoutes() {
	router := gin.Default()

	order := router.Group("/order")
	{
		order.POST("/", addOrder)
		order.GET("/:id", getOrder)
		order.DELETE("/:id", deleteOrder)
	}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

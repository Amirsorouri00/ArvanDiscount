package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	// controller "github.com/amirsorouri00/arvandiscount/controller"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcomeToDiscount)
	
	
	// Gift APIs
	
	router.NoRoute(notFound)
}

func welcomeToDiscount (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Welcome to abrarvan discount service.",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
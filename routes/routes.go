package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	controller "github.com/amirsorouri00/arvandiscount/controller"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcomeToDiscount)
	
	// Discount APIs
	router.GET("/alldiscounts", controller.GetAllDiscounts)
	router.POST("/adddiscount", controller.AddDiscount)
	// router.POST("/getdiscount", controller.GetDiscount)

	// Stream APIs (Just for test)
	router.GET("/allstreams", controller.GetAllStreams)
	router.POST("/addstream", controller.AddStream)

	// Gift APIs
	router.GET("/allgifts", controller.GetAllGifts)
	router.POST("/addgift", controller.AddGift)
	router.POST("/getgift", controller.GetGift)
	
	router.NoRoute(notFound)
}

// Welcome Test API OK
func welcomeToDiscount (c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"message": "Welcome to abrarvan discount service.",
	})
	return
}

// Not Found API OK
func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
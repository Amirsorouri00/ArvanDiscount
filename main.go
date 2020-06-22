package main

import (
	"log"
	"github.com/gin-gonic/gin"

	db	"github.com/amirsorouri00/arvandiscount/db"
	routes "github.com/amirsorouri00/arvandiscount/routes"
)

func main() {
	// Connect to DB
	db.ConnectDB()

	// Initialize Router
	router := gin.Default()

	// Route Handlers / Endpoints
	routes.Routes(router)

	log.Fatal(router.Run(":8002"))
}
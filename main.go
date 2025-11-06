package main

import (
	"log"
	"net/http"
	"otochope/database"
	"otochope/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Starting OTOCHOPE application...")

	database.Init()
	defer database.Close()

	r := gin.Default()

	config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:4200"}
	config.AllowAllOrigins = true

	r.Use(cors.New(config))
	routes.Init(r)

	if err := http.ListenAndServe(":9090", r); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}

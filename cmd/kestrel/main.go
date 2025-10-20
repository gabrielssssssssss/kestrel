package main

import (
	"github.com/gabrielssssssssss/kestrel/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	app := gin.Default()
	app.GET("/company", handlers.NewCompanyHandler().HandleCompanyRequest)
	app.GET("/mappy-search", handlers.NewMappyHandler().HandleMappySearchRequest)
	app.GET("/mappy-geo", handlers.NewMappyHandler().HandleMappyGeoRequest)
	app.Run(":8080")
}

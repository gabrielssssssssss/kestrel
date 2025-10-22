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
	app.GET("/mappy", handlers.NewMappyHandler().HandleMappy)
	app.GET("/linkedin", handlers.NewLinkedinHandler().HandleLinkedin)
	app.GET("/website", handlers.NewWebsiteHandler().HandleWebsite)
	app.Run(":8081")
}

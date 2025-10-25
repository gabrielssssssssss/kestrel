package main

import (
	handlers "github.com/gabrielssssssssss/kestrel/internal/handlers/companies"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	app := gin.Default()
	app.GET("/companies/search", handlers.NewCompaniesHandler().GetCompanyHandler)
	app.GET("/companies/sirene", handlers.NewCompaniesHandler().GetSireneHandler)
	app.Run(":8081")
}

package main

import (
	"github.com/gabrielssssssssss/kestrel/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	handler := handlers.NewCompanyHandler()
	app.GET("/company", handler.HandleCompanyRequest)
	app.Run(":8080")
}

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-shoe-store-api/database"
	"github.com/wattam/golang-shoe-store-api/handlers/shoe_handlers"
)

func main() {

	database.ConnectDatabase()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	shoes := r.Group("shoes")
	{
		shoes.GET("/get", shoe_handlers.GetAll)
		shoes.GET("/:id", shoe_handlers.Get)
		shoes.POST("/post", shoe_handlers.Post)
		shoes.PUT("/put", shoe_handlers.Put)
		shoes.DELETE("/:id", shoe_handlers.Delete)
	}

	r.Run()

	defer database.DisconnectDatabase()
}

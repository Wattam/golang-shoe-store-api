package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-shoe-store-api/controller"
	"github.com/wattam/golang-shoe-store-api/database"
)

func main() {

	database.ConnectDatabase()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		shoes := v1.Group("shoes")
		{
			shoes.GET("", controller.Index)
			shoes.GET("/:id", controller.Show)
			shoes.POST("", controller.Store)
			shoes.PUT("/:id", controller.Update)
			shoes.DELETE("/:id", controller.Delete)
		}
	}

	r.Run()

	defer database.DisconnectDatabase()
}

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
			shoes.GET("/get", controller.GetAllShoes)
			shoes.GET("/:id", controller.GetShoe)
			shoes.POST("/post", controller.PostShoe)
			shoes.PUT("/put", controller.PutShoe)
			shoes.DELETE("/:id", controller.DeleteShoe)
		}
	}

	r.Run()

	defer database.DisconnectDatabase()
}

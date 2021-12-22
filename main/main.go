package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/wattam/shoe-store-API/database"
	"github.com/wattam/shoe-store-API/handlers/shoeHandlers"
)

func main() {

	database.CreateDatabase()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	shoes := r.Group("shoes")
	{
		shoes.GET("/get", shoeHandlers.GetAllShoes)
		shoes.GET("/:id", shoeHandlers.GetShoe)
		shoes.POST("/post", shoeHandlers.CreateShoe)
		shoes.PUT("/put", shoeHandlers.EditShoe)
		shoes.DELETE("/:id", shoeHandlers.DeleteShoe)
	}

	r.Run()

	defer database.CloseConnection()
}

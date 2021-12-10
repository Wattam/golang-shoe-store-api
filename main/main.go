package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wattam/shoeStoreDB/handler"
	shoe "github.com/wattam/shoeStoreDB/models"
)

func main() {
	store := shoe.New()

	r := gin.Default()

	r.GET("/get", handler.ShoeGetAll(store))
	r.GET("/:id", handler.ShoeGet(store))
	r.POST("/post", handler.ShoeAdd(store))
	r.PUT("/put", handler.ShoeUpdate(store))
	r.DELETE("/:id", handler.ShoeDelete(store))

	r.Run()
}

package shoeHandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/shoe-store-API/database"
	"github.com/wattam/shoe-store-API/models"
)

func CreateShoe(c *gin.Context) {

	db := database.GetDatabase()

	shoe := models.Shoe{}

	c.ShouldBindJSON(&shoe)

	db.Create(&shoe)

	c.JSON(http.StatusOK, shoe)
}

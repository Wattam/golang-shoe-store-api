package shoeHandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/shoeStoreDB/database"
	"github.com/wattam/shoeStoreDB/models"
)

func GetShoe(c *gin.Context) {

	db := database.GetDatabase()

	id := c.Param("id")

	shoe := models.Shoe{}

	err := db.First(&shoe, id).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, shoe)
}
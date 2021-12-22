package shoeHandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/shoe-store-API/database"
	"github.com/wattam/shoe-store-API/models"
)

func EditShoe(c *gin.Context) {

	db := database.GetDatabase()

	shoe := models.Shoe{}

	c.ShouldBindJSON(&shoe)

	if db.First(&models.Shoe{}, shoe.ID).Error != nil {
		c.Status(http.StatusNoContent)
		return
	}

	db.Save(&shoe)

	c.JSON(http.StatusOK, shoe)
}

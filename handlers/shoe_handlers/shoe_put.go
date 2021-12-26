package shoe_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-shoe-store-api/database"
	"github.com/wattam/golang-shoe-store-api/models"
)

func Put(c *gin.Context) {

	shoe := models.Shoe{}

	c.ShouldBindJSON(&shoe)

	if database.Db.First(&models.Shoe{}, shoe.ID).Error != nil {
		c.Status(http.StatusNoContent)
		return
	}

	database.Db.Save(&shoe)

	c.JSON(http.StatusOK, shoe)
}

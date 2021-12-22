package shoeHandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/shoe-store-API/database"
	"github.com/wattam/shoe-store-API/models"
)

func GetAllShoes(c *gin.Context) {

	db := database.GetDatabase()

	shoes := []models.Shoe{}

	db.Find(&shoes)

	if len(shoes) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, shoes)
}

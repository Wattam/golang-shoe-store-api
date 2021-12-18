package shoeHandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/shoeStoreDB/database"
	"github.com/wattam/shoeStoreDB/models"
)

func DeleteShoe(c *gin.Context) {

	db := database.GetDatabase()

	id := c.Param("id")

	db.Delete(&models.Shoe{}, id)

	c.Status(http.StatusNoContent)
}

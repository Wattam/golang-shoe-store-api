package shoe_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-shoe-store-api/database"
	"github.com/wattam/golang-shoe-store-api/models"
)

func GetAll(c *gin.Context) {

	shoes := []models.Shoe{}

	database.Db.Find(&shoes)

	if len(shoes) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, shoes)
}

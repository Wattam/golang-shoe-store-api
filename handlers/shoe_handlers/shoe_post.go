package shoe_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-shoe-store-api/database"
	"github.com/wattam/golang-shoe-store-api/models"
)

func Post(c *gin.Context) {

	shoe := models.Shoe{}

	c.ShouldBindJSON(&shoe)

	database.Db.Create(&shoe)

	c.JSON(http.StatusOK, shoe)
}

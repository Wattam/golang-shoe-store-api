package shoe_handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-shoe-store-api/database"
	"github.com/wattam/golang-shoe-store-api/models"
)

func Delete(c *gin.Context) {

	id := c.Param("id")

	database.Db.Delete(&models.Shoe{}, id)

	c.Status(http.StatusNoContent)
}

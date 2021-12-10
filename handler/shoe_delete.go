package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	shoe "github.com/wattam/shoeStoreDB/models"
)

func ShoeDelete(store shoe.Deleter) gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.Param("id")
		id_converted, _ := strconv.Atoi(id)
		store.DeleteShoe(uint(id_converted))

		c.Status(http.StatusNoContent)
	}
}

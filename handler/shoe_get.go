package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	shoe "github.com/wattam/shoeStoreDB/models"
)

func ShoeGet(store shoe.Getter) gin.HandlerFunc {

	return func(c *gin.Context) {
		id := c.Param("id")
		id_converted, _ := strconv.Atoi(id)
		results := store.GetShoe(uint(id_converted))

		c.JSON(http.StatusOK, results)
	}
}

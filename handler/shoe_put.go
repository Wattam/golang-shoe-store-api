package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	shoe "github.com/wattam/shoeStoreDB/models"
)

func ShoeUpdate(store shoe.Updater) gin.HandlerFunc {

	return func(c *gin.Context) {
		request := shoePostRequest{}
		c.Bind(&request)

		id_converted, _ := strconv.Atoi(request.ID)
		price_converted, _ := strconv.ParseFloat(request.Price, 64)
		shoe := shoe.Shoe{
			ID:       uint(id_converted),
			Name:     request.Name,
			Style:    request.Style,
			Color:    request.Color,
			Material: request.Material,
			Price:    float32(price_converted),
		}

		store.UpdateShoe(shoe)

		c.Status(http.StatusNoContent)
	}
}

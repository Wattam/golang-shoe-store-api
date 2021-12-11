package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	shoe "github.com/wattam/shoeStoreDB/models"
)

type shoePostRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Style    string `json:"style"`
	Color    string `json:"color"`
	Material string `json:"material"`
	Price    string `json:"price"`
}

func ShoeAdd(store shoe.Adder) gin.HandlerFunc {

	return func(c *gin.Context) {
		request := shoePostRequest{}
		c.Bind(&request)

		price_converted, _ := strconv.ParseFloat(request.Price, 64)
		shoe := shoe.Shoe{
			ID:       0,
			Name:     request.Name,
			Style:    request.Style,
			Color:    request.Color,
			Material: request.Material,
			Price:    price_converted,
		}

		store.AddShoe(shoe)

		c.Status(http.StatusNoContent)
	}
}

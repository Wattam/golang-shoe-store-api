package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-shoe-store-api/database"
	"github.com/wattam/golang-shoe-store-api/model"
)

func GetAllShoes(c *gin.Context) {

	shoes := []model.Shoe{}

	database.Db.Find(&shoes)

	if len(shoes) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	c.JSON(http.StatusOK, shoes)
}

func GetShoe(c *gin.Context) {

	id := c.Param("id")

	shoe := model.Shoe{}

	err := database.Db.First(&shoe, id).Error

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, shoe)
}

func PostShoe(c *gin.Context) {

	shoe := model.Shoe{}

	c.ShouldBindJSON(&shoe)

	err := database.Db.Create(&shoe).Error

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, shoe)
}

func PutShoe(c *gin.Context) {

	shoe := model.Shoe{}

	c.ShouldBindJSON(&shoe)

	if database.Db.First(&model.Shoe{}, shoe.ID).Error != nil {
		c.Status(http.StatusNoContent)
		return
	}

	err := database.Db.Save(&shoe).Error

	if err != nil {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, shoe)
}

func DeleteShoe(c *gin.Context) {

	id := c.Param("id")

	database.Db.Delete(&model.Shoe{}, id)

	c.Status(http.StatusNoContent)
}

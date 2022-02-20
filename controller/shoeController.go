package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wattam/golang-shoe-store-api/database"
	"github.com/wattam/golang-shoe-store-api/model"
)

func Index(c *gin.Context) {

	shoes := []model.Shoe{}
	database.Db.Find(&shoes)

	if len(shoes) == 0 {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, shoes)
}

func Show(c *gin.Context) {

	id := c.Param("id")
	shoe := model.Shoe{}

	err := database.Db.First(&shoe, id).Error
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, shoe)
}

func Store(c *gin.Context) {

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

func Update(c *gin.Context) {

	shoe := model.Shoe{}
	c.ShouldBindJSON(&shoe)
	shoe.ID, _ = strconv.Atoi(c.Param("id"))

	if database.Db.First(&model.Shoe{}, shoe.ID).Error != nil {
		c.Status(http.StatusNotFound)
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

func Delete(c *gin.Context) {

	id := c.Param("id")

	if database.Db.First(&model.Shoe{}, id).Error != nil {
		c.Status(http.StatusNotFound)
		return
	}

	database.Db.Delete(&model.Shoe{}, id)
	c.Status(http.StatusNoContent)
}

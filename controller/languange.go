package controller

import (
	"net/http"
	"strconv"
	"testmnc/models"

	"github.com/gin-gonic/gin"
)

func AddLanguage(c *gin.Context) {
	var payload models.Languange
	if err := c.BindJSON(&payload); err != nil {
		c.String(http.StatusBadRequest, "failed")
	}

	models.AddLanguage(payload)
	c.JSON(http.StatusCreated, payload)
}

func GetOneLanguage(c *gin.Context) {
	index := c.Param("index")
	i, err := strconv.Atoi(index)
	if err != nil {
		c.String(http.StatusBadRequest, "parameter index should be a number")
		return
	}
	c.JSON(http.StatusOK, models.GetLanguage(i))
}

func GetAllLanguage(c *gin.Context) {
	c.JSON(http.StatusOK, models.GetAllLanguage())
}

func DeleteLanguage(c *gin.Context) {
	index := c.Param("index")
	i, err := strconv.Atoi(index)
	if err != nil {
		c.String(http.StatusBadRequest, "parameter index should be a number")
		return
	}

	if err := models.RemoveLanguage(i); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}
	c.String(http.StatusOK, "success removing language")
}

func UpdateLanguage(c *gin.Context) {
	index := c.Param("index")
	i, err := strconv.Atoi(index)
	if err != nil {
		c.String(http.StatusBadRequest, "parameter index should be a number")
		return
	}

	var payload models.Languange
	if err := c.BindJSON(&payload); err != nil {
		c.String(http.StatusBadRequest, "failed")
	}

	models.UpdateLanguage(i, payload)
	c.JSON(http.StatusCreated, payload)
}

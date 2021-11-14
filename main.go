package main

import (
	"net/http"
	"testmnc/controller"
	"testmnc/models"

	"github.com/gin-gonic/gin"
)

func main() {

	models.InitLanguage()

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello GO Developers")
	})
	r.GET("/palindrome/:text", controller.CheckPalindrome)

	r.POST("/language", controller.AddLanguage)
	r.GET("/language/:index", controller.GetOneLanguage)
	r.GET("/languages", controller.GetAllLanguage)
	r.PATCH("language/:index", controller.UpdateLanguage)
	r.DELETE("language/:index", controller.DeleteLanguage)

	r.Run(":8080")
}

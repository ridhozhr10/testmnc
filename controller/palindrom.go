package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CheckPalindrome(c *gin.Context) {
	text := c.Param("text")
	if isPalindrome(text) {
		c.String(http.StatusOK, "Palindrome")
		return
	}
	c.String(http.StatusBadRequest, "Not palindrome")
}

func isPalindrome(value string) bool {
	for i := 0; i < len(value); i++ {
		if value[i] != value[len(value)-i-1] {
			return false
		}
	}

	return true
}

package controller

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.True(t, isPalindrome("katak"))
	assert.False(t, isPalindrome("lele"))
}

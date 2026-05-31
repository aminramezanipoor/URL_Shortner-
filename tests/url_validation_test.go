package tests

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsValidURL(url string) bool {
	return strings.HasPrefix(url, "http://") ||
		strings.HasPrefix(url, "https://")
}

func TestValidURL(t *testing.T) {
	assert.True(t, IsValidURL("https://quera.org"))
	assert.True(t, IsValidURL("http://google.com"))
}

func TestInvalidURL(t *testing.T) {
	assert.False(t, IsValidURL("google.com"))
	assert.False(t, IsValidURL(""))
	assert.False(t, IsValidURL("test"))
}
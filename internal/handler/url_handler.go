package handler

import (
	"net/http"

	"github.com/aminramezanipoor/url-shortener/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type URLHandler struct {
	urlService *service.URLService
}

func NewURLHandler(urlService *service.URLService) *URLHandler {
	return &URLHandler{urlService: urlService}
}

type CreateURLRequest struct {
	OriginalURL string `json:"original_url" binding:"required"`
}

func (h *URLHandler) CreateShortURL(c *gin.Context) {
	var req CreateURLRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}

	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	userID, ok := userIDValue.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user id",
		})
		return
	}

	createdURL, err := h.urlService.CreateShortURL(req.OriginalURL, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	shortURL := "http://localhost:8080/" + createdURL.ShortCode

	c.JSON(http.StatusCreated, gin.H{
		"message":      "short url created successfully",
		"original_url": createdURL.OriginalURL,
		"short_code":   createdURL.ShortCode,
		"short_url":    shortURL,
	})
}

func (h *URLHandler) GetMyURLs(c *gin.Context) {
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	userID, ok := userIDValue.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "invalid user id",
		})
		return
	}

	urls, err := h.urlService.GetUserURLs(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not get urls",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"urls": urls,
	})
}

func (h *URLHandler) Redirect(c *gin.Context) {
	shortCode := c.Param("shortCode")

	foundURL, err := h.urlService.GetOriginalURL(shortCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "url not found",
		})
		return
	}

	c.Redirect(http.StatusFound, foundURL.OriginalURL)
}
func (h *URLHandler) GetAllURLs(c *gin.Context) {
	urls, err := h.urlService.GetAllURLs()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not get urls",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"urls": urls,
	})
}

func (h *URLHandler) DeleteURL(c *gin.Context) {
	idParam := c.Param("id")

	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid url id",
		})
		return
	}

	if err := h.urlService.DeleteURL(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not delete url",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "url deleted successfully",
	})
}
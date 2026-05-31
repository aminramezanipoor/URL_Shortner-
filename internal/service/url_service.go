package service

import (
	"errors"
	"net/url"

	"github.com/aminramezanipoor/url-shortener/internal/models"
	"github.com/aminramezanipoor/url-shortener/internal/repository"
	"github.com/aminramezanipoor/url-shortener/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type URLService struct {
	urlRepo *repository.URLRepository
}

func NewURLService(urlRepo *repository.URLRepository) *URLService {
	return &URLService{urlRepo: urlRepo}
}

func (s *URLService) CreateShortURL(originalURL string, userID uuid.UUID) (*models.URL, error) {
	if !isValidURL(originalURL) {
		return nil, errors.New("invalid url")
	}

	var shortCode string

	for i := 0; i < 5; i++ {
		code, err := utils.GenerateShortCode(7)
		if err != nil {
			return nil, err
		}

		_, err = s.urlRepo.FindByShortCode(code)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			shortCode = code
			break
		}

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	if shortCode == "" {
		return nil, errors.New("could not generate unique short code")
	}

	newURL := &models.URL{
		OriginalURL: originalURL,
		ShortCode:   shortCode,
		UserID:      userID,
	}

	if err := s.urlRepo.Create(newURL); err != nil {
		return nil, err
	}

	return newURL, nil
}

func (s *URLService) GetUserURLs(userID uuid.UUID) ([]models.URL, error) {
	return s.urlRepo.FindByUserID(userID)
}

func (s *URLService) GetOriginalURL(shortCode string) (*models.URL, error) {
	foundURL, err := s.urlRepo.FindByShortCode(shortCode)
	if err != nil {
		return nil, errors.New("url not found")
	}

	if err := s.urlRepo.IncrementClicks(foundURL.ID); err != nil {
		return nil, err
	}

	return foundURL, nil
}

func isValidURL(rawURL string) bool {
	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return false
	}

	return parsedURL.Scheme == "http" || parsedURL.Scheme == "https"
}
func (s *URLService) GetAllURLs() ([]models.URL, error) {
	return s.urlRepo.FindAll()
}

func (s *URLService) DeleteURL(id uuid.UUID) error {
	return s.urlRepo.DeleteByID(id)
}
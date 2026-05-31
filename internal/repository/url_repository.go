package repository

import (
	"github.com/aminramezanipoor/url-shortener/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type URLRepository struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) *URLRepository {
	return &URLRepository{db: db}
}

func (r *URLRepository) Create(url *models.URL) error {
	return r.db.Create(url).Error
}

func (r *URLRepository) FindByShortCode(shortCode string) (*models.URL, error) {
	var url models.URL

	err := r.db.Where("short_code = ?", shortCode).First(&url).Error
	if err != nil {
		return nil, err
	}

	return &url, nil
}

func (r *URLRepository) FindByUserID(userID uuid.UUID) ([]models.URL, error) {
	var urls []models.URL

	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&urls).Error
	return urls, err
}

func (r *URLRepository) IncrementClicks(id uuid.UUID) error {
	return r.db.Model(&models.URL{}).
		Where("id = ?", id).
		UpdateColumn("clicks", gorm.Expr("clicks + ?", 1)).
		Error
}
func (r *URLRepository) FindAll() ([]models.URL, error) {
	var urls []models.URL

	err := r.db.Order("created_at DESC").Find(&urls).Error
	return urls, err
}

func (r *URLRepository) DeleteByID(id uuid.UUID) error {
	return r.db.Delete(&models.URL{}, "id = ?", id).Error
}
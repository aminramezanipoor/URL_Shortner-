package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type URL struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	OriginalURL string    `gorm:"not null"`
	ShortCode   string    `gorm:"unique;not null"`
	Clicks      int64     `gorm:"default:0"`

	UserID uuid.UUID

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *URL) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.New()
	return nil
}
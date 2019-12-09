package entity

import (
	"bookstore/db/model/hooks"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Book db entity representation of a user
type Book struct {
	hooks.BeforeCreation
	ID            uuid.UUID       `gorm:"type:varchar(36);primary_key;not null;"`
	ThumbnailURL  string          `gorm:"type:varchar(255);not null;"`
	Status        string          `gorm:"type:varchar(10);not null;"`
	ImagePath     string          `gorm:"type:varchar(255);not null;"`
	Title         string          `gorm:"type:varchar(100);not null;"`
	PublishedDate *time.Time      `gorm:"not null;"`
	Authors       []Author        `gorm:"many2many:author_book"`
	Price         float64         `gorm:"not null;"`
	BookingCost   float64         `gorm:"not null;"`
	Booked        json.RawMessage `gorm:"not null;"`
	Description   string          `gorm:"type:text;not null;"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     *time.Time `sql:"index"`
}

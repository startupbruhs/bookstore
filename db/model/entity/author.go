package entity

import (
	"bookstore/db/model/hooks"
	"time"

	"github.com/google/uuid"
)

// Author db entity representation of a user
type Author struct {
	hooks.BeforeCreation
	ID        uuid.UUID `gorm:"type:varchar(36);primary_key;not null;"`
	Name      string    `gorm:"type:varchar(50);not null"`
	Books     []*Book   `gorm:"many2many:author_book;not null;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

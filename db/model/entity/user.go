package entity

import (
	"startupbruhs.github.io/bookstore/db/model/hooks"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// User db entity representation of a user
type User struct {
	hooks.BeforeCreation
	ID              uuid.UUID `gorm:"type:varchar(36);primary_key;not null;"`
	RoleID          sql.NullInt64
	Name            string    `gorm:"varchar(50);not null;"`
	Email           string    `gorm:"varchar(255);not null;unique;"`
	Avatar          string    `gorm:"varchar(255);not null;"`
	EmailVerifiedAt time.Time `gorm:"default:null;"`
	Password        string    `gorm:"varchar(64);not null;"`
	RememberToken   string    `gorm:"varchar(64);"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       *time.Time `sql:"index"`
}

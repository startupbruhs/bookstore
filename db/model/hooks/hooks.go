package hooks

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type BeforeCreation struct {
}

func (toBeCreated *BeforeCreation) BeforeCreate(scope *gorm.Scope) error {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return err
	}
	return scope.SetColumn("ID", uuid)
}

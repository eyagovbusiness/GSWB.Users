package models

import (
	"time"

	"github.com/google/uuid"
)

type UserModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);not null;uniqueIndex"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

// TableName allows full control over the DB table name
func (UserModel) TableName() string {
	return "users"
}

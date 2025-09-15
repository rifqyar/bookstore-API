package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint            `gorm:"primaryKey" json:"id"`
	Name      string          `gorm:"size:100;not null" json:"name"`
	Email     string          `gorm:"size:100;uniqueIndex;not null" json:"email"`
	Password  string          `gorm:"size:255;not null" json:"-"`
	Role      string          `gorm:"type:VARCHAR(10) CHECK (role IN ('user','admin'));default:'user'" json:"role"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
	IsActive  bool            `gorm:"default:true" json:"is_active"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

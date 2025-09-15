package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"size:100;not null;unique" json:"name"`
}

type Book struct {
	ID          uint            `gorm:"primaryKey" json:"id"`
	Title       string          `gorm:"size:255;not null" json:"title"`
	Author      string          `gorm:"size:100;not null" json:"author"`
	Price       float64         `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock       int             `gorm:"not null" json:"stock"`
	Year        int             `json:"year"`
	CategoryID  uint            `json:"category_id"`
	Category    Category        `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	ImageBase64 string          `gorm:"type:text" json:"image_base64"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   *gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

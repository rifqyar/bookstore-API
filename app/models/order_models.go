package models

import (
	"time"
)

type Order struct {
	ID         uint        `gorm:"primaryKey" json:"id"`
	UserID     uint        `json:"user_id"`
	User       User        `gorm:"foreignKey:UserID" json:"user,omitempty"`
	TotalPrice float64     `gorm:"type:decimal(10,2)" json:"total_price"`
	Status     string      `gorm:"type:VARCHAR(20);default:'PENDING'" json:"status"`
	CreatedAt  time.Time   `json:"created_at"`
	Items      []OrderItem `json:"items" gorm:"constraint:OnDelete:CASCADE"`
}

type OrderItem struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	OrderID  uint    `json:"order_id"`
	BookID   uint    `json:"book_id"`
	Book     Book    `gorm:"foreignKey:BookID" json:"book,omitempty"`
	Quantity int     `json:"quantity"`
	Price    float64 `gorm:"type:decimal(10,2)" json:"price"`
}

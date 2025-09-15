package db

import (
	"bookstore-api/app/config"
	"bookstore-api/app/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPass, cfg.DBName, cfg.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Migrating Database...")
	if err := db.AutoMigrate(
		&models.User{},
		&models.Category{},
		&models.Book{},
		&models.Order{},
		&models.OrderItem{},
	); err != nil {
		log.Fatalf("Failed Migrating Database: %v", err)
		return nil, err
	}
	return db, nil
}

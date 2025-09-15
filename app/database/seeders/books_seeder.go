package seeders

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"bookstore-api/app/models"

	"github.com/bxcodec/faker/v4"
	"gorm.io/gorm"
)

func BookSeeder(db *gorm.DB) {
	categories := []string{"Fiksi", "Non-Fiksi", "Teknologi", "Sejarah", "Sains", "Bisnis"}

	for _, name := range categories {
		var count int64
		db.Model(&models.Category{}).Where("name = ?", name).Count(&count)
		if count == 0 {
			db.Create(&models.Category{Name: name})
			fmt.Println("Default category created: ", name)
		}
	}

	var allCategories []models.Category
	db.Find(&allCategories)

	for i := 0; i < 20; i++ {
		category := allCategories[rand.Intn(len(allCategories))]

		image := []byte(fmt.Sprintf("DummyImage-%d", i))
		imageBase64 := base64.StdEncoding.EncodeToString(image)

		book := models.Book{
			Title:       faker.Sentence(),
			Author:      faker.Name(),
			Price:       float64(rand.Intn(500000))/100.0 + 10,
			Stock:       rand.Intn(100) + 1,
			Year:        rand.Intn(30) + 1990,
			CategoryID:  category.ID,
			ImageBase64: imageBase64,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		db.Create(&book)
		fmt.Println("Book seeder created: ", book.Title+" in category "+category.Name)
	}
}

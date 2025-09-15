package seeders

import (
	"fmt"
	"log"

	"bookstore-api/app/models"
	"bookstore-api/app/utils"

	"github.com/bxcodec/faker/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func UserSeeder(db *gorm.DB) {
	adminPass, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	admin := models.User{
		Name:     "Administrator",
		Email:    "admin.bookstore@mail.com",
		Password: string(adminPass),
		Role:     "admin",
		IsActive: true,
	}
	if err := db.Create(&admin).Error; err != nil {
		log.Println("Admin sudah ada atau gagal dibuat:", err)
	}

	userPass, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	user := models.User{
		Name:     "John Doe",
		Email:    "john_doe@mail.com",
		Password: string(userPass),
		Role:     "user",
		IsActive: true,
	}

	if err := db.Create(&user).Error; err != nil {
		log.Println("User Default sudah ada atau gagal dibuat:", err)
	}

	for i := 0; i < 10; i++ {
		hashedPass, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		user := models.User{
			Name:     faker.Name(),
			Email:    utils.GenerateEmailFromName(faker.Name()),
			Password: string(hashedPass),
			Role:     "user",
			IsActive: true,
		}
		db.Create(&user)
		fmt.Println("User faker dibuat:", user.Email)
	}
}

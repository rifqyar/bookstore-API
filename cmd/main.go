package main

import (
	"bookstore-api/app/config"
	"bookstore-api/app/database/seeders"
	"bookstore-api/app/db"
	"bookstore-api/app/routes"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	docs "bookstore-api/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Bookstore REST API
// @version 1.0
// @description This is a RESTful API for a bookstore application (technical test).
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email rifqyaditya55@gmail.com

// @host localhost:8080
// @BasePath /
// @schemes http

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg := config.Load()
	gormDB, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "seed:db":
			fmt.Println("Running Seeders...")
			seeders.UserSeeder(gormDB)
			seeders.BookSeeder(gormDB)
			fmt.Println("Database Ready")
		default:
			fmt.Println("Command not found")
		}
	}

	docs.SwaggerInfo.BasePath = "/"
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.RegisterRoutes(r, gormDB, cfg)

	addr := ":" + cfg.AppPort
	log.Println("listening on", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}
}

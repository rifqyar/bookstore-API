package routes

import (
	"bookstore-api/app/config"
	"bookstore-api/app/handlers"
	"bookstore-api/app/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB, cfg *config.Config) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Bookstore API V1.0",
		})
	})

	r.POST("/register", handlers.Register(db))
	r.POST("/login", handlers.Login(db, cfg))

	auth := r.Group("/")
	auth.Use(middleware.JWTAuth(cfg))
	{
		cat := auth.Group("/categories")
		cat.GET("", handlers.ListCategories(db))
		cat.GET("/:id", func(c *gin.Context) {})
		cat.POST("", middleware.RequireRole("admin"), handlers.CreateCategory(db))
		cat.PUT("/:id", middleware.RequireRole("admin"), handlers.UpdateCategory(db))
		cat.DELETE("/:id", middleware.RequireRole("admin"), handlers.DeleteCategory(db))

		book := auth.Group("/books")
		book.GET("", handlers.ListBooks(db))
		book.GET("/:id", handlers.GetBook(db))
		book.POST("", middleware.RequireRole("admin"), handlers.CreateBook(db))
		book.PUT("/:id", middleware.RequireRole("admin"), handlers.UpdateBook(db))
		book.DELETE("/:id", middleware.RequireRole("admin"), handlers.DeleteBook(db))

		orders := auth.Group("/orders")
		orders.POST("", handlers.CreateOrder(db))
		orders.POST("/:id/pay", handlers.PayOrder(db))
		orders.GET("", handlers.ListOrders(db))
		orders.GET("/:id", handlers.GetOrder(db))

		reports := auth.Group("/reports")
		reports.Use(middleware.RequireRole("admin"))
		reports.GET("/sales", handlers.SalesReport(db))
		reports.GET("/bestseller", handlers.BestsellerReport(db))
		reports.GET("/prices", handlers.PriceStatsReport(db))
	}
}

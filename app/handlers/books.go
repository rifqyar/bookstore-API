package handlers

import (
	"net/http"
	"strconv"

	"bookstore-api/app/dto"
	"bookstore-api/app/models"
	"bookstore-api/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateBook godoc
// @Summary Create book
// @Tags Books
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.CreateBook true "Book info"
// @Success 201 {object} map[string]interface{}
// @Router /books [post]
func CreateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.CreateBook
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.JSONError(c, http.StatusBadRequest, err.Error())
			return
		}
		book := models.Book{
			Title: req.Title, Author: req.Author, Price: req.Price,
			Stock: req.Stock, Year: req.Year, CategoryID: req.CategoryID, ImageBase64: req.ImageBase64,
		}
		if err := db.Create(&book).Error; err != nil {
			utils.JSONError(c, http.StatusBadRequest, err.Error())
			return
		}
		utils.JSONCreated(c, "Success Created Book Data", book)
	}
}

// ListBooks godoc
// @Summary List books
// @Tags Books
// @Security BearerAuth
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param q query string false "Search keyword (title or author)"
// @Param category query string false "Filter by category id or name"
// @Success 200 {object} map[string]interface{}
// @Router /books [get]
func ListBooks(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// pagination & search & filter
		pageStr := c.DefaultQuery("page", "1")
		limitStr := c.DefaultQuery("limit", "10")
		q := c.Query("q") // title or author
		cat := c.Query("category")

		page, _ := strconv.Atoi(pageStr)
		limit, _ := strconv.Atoi(limitStr)
		if page < 1 {
			page = 1
		}
		offset := (page - 1) * limit

		var books []models.Book
		query := db.Preload("Category").Model(&models.Book{})
		if q != "" {
			like := "%" + q + "%"
			query = query.Where("title ILIKE ? OR author ILIKE ?", like, like)
		}
		if cat != "" {
			// allow category id or name
			if id, err := strconv.Atoi(cat); err == nil {
				query = query.Where("category_id = ?", id)
			} else {
				// join category by name
				query = query.Joins("JOIN categories ON categories.id = books.category_id").Where("categories.name ILIKE ?", "%"+cat+"%")
			}
		}
		var total int64
		query.Count(&total)
		query = query.Limit(limit).Offset(offset).Order("id desc").Find(&books)
		utils.JSONOk(c, gin.H{"items": books, "page": page, "limit": limit, "total": total})
	}
}

// GetBook godoc
// @Summary Get book
// @Tags Books
// @Security BearerAuth
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]interface{}
// @Router /books/{id} [get]
func GetBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var book models.Book
		if err := db.Preload("Category").First(&book, id).Error; err != nil {
			utils.JSONError(c, http.StatusNotFound, "book not found")
			return
		}
		utils.JSONOk(c, book)
	}
}

// UpdateBook godoc
// @Summary Update book
// @Tags Books
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param request body dto.UpdateBook true "Book info"
// @Success 200 {object} map[string]interface{}
// @Router /books/{id} [put]
func UpdateBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var book models.Book
		if err := db.First(&book, id).Error; err != nil {
			utils.JSONError(c, http.StatusNotFound, "book not found")
			return
		}

		var req dto.UpdateBook
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.JSONError(c, http.StatusBadRequest, err.Error())
			return
		}

		updates := map[string]interface{}{}
		if req.Title != nil {
			updates["title"] = *req.Title
		}
		if req.Author != nil {
			updates["author"] = *req.Author
		}
		if req.Price != nil {
			updates["price"] = *req.Price
		}
		if req.Stock != nil {
			updates["stock"] = *req.Stock
		}
		if req.Year != nil {
			updates["year"] = *req.Year
		}
		if req.CategoryID != nil {
			updates["category_id"] = *req.CategoryID
		}
		if req.ImageBase64 != nil {
			updates["image_base64"] = *req.ImageBase64
		}

		if len(updates) > 0 {
			if err := db.Model(&book).Updates(updates).Error; err != nil {
				utils.JSONError(c, http.StatusInternalServerError, "failed to update book: "+err.Error())
				return
			}
		}

		utils.JSONOk(c, book)
	}

}

// DeleteBook godoc
// @Summary Delete book
// @Tags Books
// @Security BearerAuth
// @Param id path int true "Book ID"
// @Success 200 {object} map[string]interface{}
// @Router /books/{id} [delete]
func DeleteBook(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var book models.Book
		if err := db.First(&book, id).Error; err != nil {
			utils.JSONError(c, http.StatusNotFound, "book not found")
			return
		}

		if err := db.Delete(&models.Book{}, id).Error; err != nil {
			utils.JSONError(c, http.StatusBadRequest, err.Error())
			return
		}
		utils.JSONOk(c, gin.H{"message": "deleted"})
	}
}

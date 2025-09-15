package handlers

import (
	"bookstore-api/app/dto"
	"bookstore-api/app/models"
	"bookstore-api/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateCategory godoc
// @Summary Create category
// @Description Admin creates a new category
// @Tags Categories
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.CategoryRequest true "Category info"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /categories [post]
func CreateCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.CategoryRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.JSONError(c, http.StatusBadRequest, err.Error())
			return
		}
		cat := models.Category{Name: req.Name}
		if err := db.Create(&cat).Error; err != nil {
			utils.JSONError(c, http.StatusBadRequest, err.Error())
			return
		}
		utils.JSONCreated(c, "Success created category", cat)
	}
}

// ListCategories godoc
// @Summary List categories
// @Tags Categories
// @Security BearerAuth
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /categories [get]
func ListCategories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var cats []models.Category
		db.Find(&cats)
		utils.JSONOk(c, cats)
	}
}

// UpdateCategory godoc
// @Summary Update books category
// @Tags Categories
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Category ID"
// @Param request body dto.CategoryRequest true "Category info"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /categories/{id} [put]
func UpdateCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var cat models.Category
		id := c.Param("id")
		if err := db.First(&cat, id).Error; err != nil {
			utils.JSONError(c, http.StatusNotFound, "category not found")
			return
		}
		var req dto.CategoryRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.JSONError(c, http.StatusBadRequest, err.Error())
			return
		}
		cat.Name = req.Name
		db.Save(&cat)
		utils.JSONOk(c, cat)
	}
}

// DeleteCategory godoc
// @Summary Delete category
// @Tags Categories
// @Security BearerAuth
// @Param id path int true "Category ID"
// @Success 200 {object} map[string]interface{}
// @Router /categories/{id} [delete]
func DeleteCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var cat models.Category
		id := c.Param("id")
		if err := db.First(&cat, id).Error; err != nil {
			utils.JSONError(c, http.StatusNotFound, "category not found")
			return
		}
		db.Delete(&cat)
		utils.JSONOk(c, gin.H{"message": "deleted"})
	}
}

package handlers

import (
	"bookstore-api/app/dto"
	"bookstore-api/app/models"
	"bookstore-api/app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateOrder godoc
// @Summary Create order
// @Tags Orders
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body dto.CreateOrderRequest true "Order items"
// @Success 201 {object} map[string]interface{}
// @Router /orders [post]
func CreateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.CreateOrderRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.JSONError(c, http.StatusBadRequest, err.Error())
			return
		}

		userIDv, _ := c.Get("user_id")
		userID := userIDv.(uint)

		tx := db.Begin()
		if tx.Error != nil {
			utils.JSONError(c, http.StatusInternalServerError, "could not start tx")
			return
		}

		order := models.Order{UserID: userID, Status: "PENDING"}
		if err := tx.Create(&order).Error; err != nil {
			tx.Rollback()
			utils.JSONError(c, http.StatusInternalServerError, err.Error())
			return
		}

		total := 0.0
		for _, it := range req.Items {
			var book models.Book
			if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&book, it.BookID).Error; err != nil {
				tx.Rollback()
				utils.JSONError(c, http.StatusBadRequest, "book not found")
				return
			}
			if it.Quantity > book.Stock {
				tx.Rollback()
				utils.JSONError(c, http.StatusBadRequest, "quantity exceeds stock for book "+book.Title)
				return
			}

			book.Stock -= it.Quantity
			if err := tx.Save(&book).Error; err != nil {
				tx.Rollback()
				utils.JSONError(c, http.StatusInternalServerError, err.Error())
				return
			}

			price := book.Price * float64(it.Quantity)
			oi := models.OrderItem{
				OrderID: order.ID, BookID: book.ID, Quantity: it.Quantity, Price: book.Price,
			}
			if err := tx.Create(&oi).Error; err != nil {
				tx.Rollback()
				utils.JSONError(c, http.StatusInternalServerError, err.Error())
				return
			}
			total += price
		}
		order.TotalPrice = total
		if err := tx.Save(&order).Error; err != nil {
			tx.Rollback()
			utils.JSONError(c, http.StatusInternalServerError, err.Error())
			return
		}
		tx.Commit()

		if err := db.Preload("User").Preload("Items.Book.Category").Preload("Items.Book").First(&order, order.ID).Error; err != nil {
			utils.JSONError(c, http.StatusInternalServerError, "could not fetch order")
			return
		}

		utils.JSONCreated(c, "Success Order Book", order)
	}
}

// PayOrder godoc
// @Summary Pay order
// @Tags Orders
// @Security BearerAuth
// @Param id path int true "Order ID"
// @Success 200 {object} map[string]interface{}
// @Router /orders/{id}/pay [post]
func PayOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		userIDv, _ := c.Get("user_id")
		userID := userIDv.(uint)
		var order models.Order
		if err := db.Preload("User").Preload("Items.Book.Category").Preload("Items.Book").First(&order, id).Error; err != nil {
			utils.JSONError(c, http.StatusNotFound, "order not found")
			return
		}
		rolev, _ := c.Get("role")
		role := rolev.(string)
		if role != "admin" && order.UserID != userID {
			utils.JSONError(c, http.StatusForbidden, "not authorized")
			return
		}
		if order.Status != "PENDING" {
			utils.JSONError(c, http.StatusBadRequest, "order has been paid or cancelled")
			return
		}

		order.Status = "PAID"
		db.Save(&order)
		utils.JSONOk(c, order)
	}
}

// ListOrders godoc
// @Summary List orders
// @Tags Orders
// @Security BearerAuth
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /orders [get]
func ListOrders(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rolev, _ := c.Get("role")
		role := rolev.(string)
		userIDv, _ := c.Get("user_id")
		userID := userIDv.(uint)

		var orders []models.Order
		q := db.Preload("Items.Book").Order("id desc")
		if role != "admin" {
			q = q.Where("user_id = ?", userID)
		}
		q.Preload("User").Preload("Items.Book.Category").Preload("Items.Book").Find(&orders)
		utils.JSONOk(c, orders)
	}
}

// GetOrder godoc
// @Summary Get order
// @Tags Orders
// @Security BearerAuth
// @Produce json
// @Param id path int true "Order ID"
// @Success 200 {object} map[string]interface{}
// @Router /orders/{id} [get]
func GetOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		rolev, _ := c.Get("role")
		role := rolev.(string)
		userIDv, _ := c.Get("user_id")
		userID := userIDv.(uint)

		var order models.Order
		if err := db.Preload("User").Preload("Items.Book.Category").Preload("Items.Book").First(&order, id).Error; err != nil {
			utils.JSONError(c, http.StatusNotFound, "order not found")
			return
		}
		if role != "admin" && order.UserID != userID {
			utils.JSONError(c, http.StatusForbidden, "not authorized")
			return
		}
		utils.JSONOk(c, order)
	}
}

package handlers

import (
	"bookstore-api/app/dto"
	"bookstore-api/app/models"
	"bookstore-api/app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SalesReport godoc
// @Summary Sales report
// @Description Show total revenue and total books sold
// @Tags Reports
// @Security BearerAuth
// @Produce json
// @Success 200 {object} dto.SalesReportResponse
// @Router /reports/sales [get]
func SalesReport(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// total omzet (sum total_price where status PAID), total books sold (sum quantity)
		var totalRevenue float64
		var totalBooksSold int64
		db.Model(&models.Order{}).Where("status = ?", "PAID").Select("COALESCE(SUM(total_price),0)").Scan(&totalRevenue)
		db.Model(&models.OrderItem{}).Joins("JOIN orders on orders.id = order_items.order_id").Where("orders.status = ?", "PAID").Select("COALESCE(SUM(order_items.quantity),0)").Scan(&totalBooksSold)
		utils.JSONOk(c, gin.H{"revenue": totalRevenue, "books_sold": totalBooksSold})
	}
}

// BestsellerReport godoc
// @Summary Bestseller report
// @Description Show top 3 best selling books
// @Tags Reports
// @Security BearerAuth
// @Produce json
// @Success 200 {array} dto.BestsellerReportResponse
// @Router /reports/bestseller [get]
func BestsellerReport(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// top 3 by sold quantity (only PAID orders)
		type Row dto.BestsellerReportResponse
		var rows []Row
		db.Raw(`
            SELECT b.id as book_id, b.title, SUM(oi.quantity) as sold
            FROM order_items oi
            JOIN orders o ON o.id = oi.order_id
            JOIN books b ON b.id = oi.book_id
            WHERE o.status = ?
            GROUP BY b.id, b.title
            ORDER BY sold DESC
            LIMIT 3
        `, "PAID").Scan(&rows)
		utils.JSONOk(c, rows)
	}
}

// PriceStatsReport godoc
// @Summary Price stats
// @Description Show max, min, and average price of books
// @Tags Reports
// @Security BearerAuth
// @Produce json
// @Success 200 {object} dto.PriceStatsReportResponse
// @Router /reports/prices [get]
func PriceStatsReport(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var max, min, avg float64
		db.Model(&models.Book{}).Select("COALESCE(MAX(price),0)").Scan(&max)
		db.Model(&models.Book{}).Select("COALESCE(MIN(price),0)").Scan(&min)
		db.Model(&models.Book{}).Select("COALESCE(AVG(price),0)").Scan(&avg)
		utils.JSONOk(c, gin.H{"max": max, "min": min, "avg": avg})
	}
}

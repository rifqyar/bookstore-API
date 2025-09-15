package handlers

import (
	"bookstore-api/app/config"
	"bookstore-api/app/dto"
	"bookstore-api/app/models"
	"bookstore-api/app/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Register godoc
// @Summary Register new user
// @Description Create a new user account
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.UserRequest true "User info"
// @Success 201 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /register [post]
func Register(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.UserRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.JSONError(c, http.StatusBadRequest, err.Error())
			return
		}
		hashed, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		user := models.User{Name: req.Name, Email: req.Email, Password: string(hashed), Role: "user"}
		if err := db.Create(&user).Error; err != nil {
			utils.JSONError(c, http.StatusBadRequest, err.Error())
			return
		}
		utils.JSONCreated(c, "User Created Successfully", gin.H{
			"id":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"role":       user.Role,
			"is_active":  user.IsActive,
			"created_at": user.CreatedAt,
		})
	}
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body dto.UserLoginRequest true "Login info"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /login [post]
func Login(db *gorm.DB, cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req dto.UserLoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.JSONError(c, http.StatusBadRequest, err.Error())
			return
		}
		var user models.User
		if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
			utils.JSONError(c, http.StatusUnauthorized, "Email not found in our database")
			return
		}
		if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
			utils.JSONError(c, http.StatusUnauthorized, "Password is incorrect")
			return
		}

		expired := time.Now().Add(72 * time.Hour)
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub":  user.ID,
			"role": user.Role,
			"exp":  expired.Unix(),
		})
		s, err := token.SignedString([]byte(cfg.JWTSecret))
		if err != nil {
			utils.JSONError(c, http.StatusInternalServerError, "could not create token")
			return
		}

		utils.JSONOk(c, gin.H{
			"success": true,
			"message": "Login successful",
			"data": gin.H{
				"access_token": s,
				"token_type":   "bearer",
				"expires_in":   int(expired.Sub(time.Now()).Seconds()),
				"user": gin.H{
					"id":    user.ID,
					"name":  user.Name,
					"email": user.Email,
					"role":  user.Role,
				},
			},
		})
	}
}

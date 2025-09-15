package utils

import "github.com/gin-gonic/gin"

func JSONOk(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{"status": "success", "data": data})
}

func JSONCreated(c *gin.Context, message string, data interface{}) {
	c.JSON(201, gin.H{"status": "success", "data": data})
}

func JSONError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"status": "error", "message": message})
}

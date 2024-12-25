package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Khởi tạo router Gin
	r := gin.Default()

	// Định nghĩa endpoint GET /
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Pong Welcome to the Gin server 1!",
		})
	})

	// Định nghĩa endpoint POST /data
	r.POST("/data", func(c *gin.Context) {
		var input struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}

		// Xử lý dữ liệu JSON từ request
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
		}

		// Trả lại phản hồi
		c.JSON(http.StatusOK, gin.H{
			"message": "Data received successfully",
			"data":    input,
		})
	})

	// Định nghĩa endpoint GET /health
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "Server is running",
		})
	})

	// Chạy server trên cổng 4080
	r.Run(":4080")
}

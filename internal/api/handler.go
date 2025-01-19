package api

import (
	"base-code-api/internal/domain"
	"base-code-api/internal/middleware"
	"base-code-api/internal/model"
	"base-code-api/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

var barangService = service.NewBarangService(domain.NewBarangRepository())
var userService = service.NewUserService()

func StartServer() {
	r := gin.Default()

	r.Use(middleware.AuthMiddleware())

	r.POST("/register", userService.Register)
	r.POST("/login", userService.Login)

	r.POST("/barang", func(c *gin.Context) {
		var barang model.Barang
		if err := c.ShouldBindJSON(&barang); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		createdBarang := barangService.CreateBarang(barang)
		c.JSON(http.StatusOK, createdBarang)
	})

	r.GET("/barang/:id", func(c *gin.Context) {
		id := c.Param("id")
		barang := barangService.GetBarang(id)
		if barang == nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Barang not found"})
			return
		}
		c.JSON(http.StatusOK, barang)
	})

	r.PUT("/barang/:id", func(c *gin.Context) {
		id := c.Param("id")
		var barang model.Barang
		if err := c.ShouldBindJSON(&barang); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		updatedBarang := barangService.UpdateBarang(id, barang)
		if updatedBarang == nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "Barang not found"})
			return
		}
		c.JSON(http.StatusOK, updatedBarang)
	})

	r.DELETE("/barang/:id", func(c *gin.Context) {
		id := c.Param("id")
		deleted := barangService.DeleteBarang(id)
		if !deleted {
			c.JSON(http.StatusNotFound, gin.H{"message": "Barang not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Barang deleted"})
	})

	r.Run(":8080")
}

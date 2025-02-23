package api

import (
	"base-code-api/app/domain"
	"base-code-api/app/middleware"
	"base-code-api/app/model"
	"base-code-api/app/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var barangService = service.NewBarangService(domain.NewBarangRepository())
var userService = service.NewUserService()

func StartServer() {
	r := gin.Default()
	APP_PORT := os.Getenv("APP_PORT")

	authorized := r.Group("/api")
	authorized.Use(middleware.AuthMiddleware())
	{
		authorized.POST("/barang", func(c *gin.Context) {
			var barang model.Barang
			if err := c.ShouldBindJSON(&barang); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			createdBarang := barangService.CreateBarang(barang)
			c.JSON(http.StatusOK, createdBarang)
		})

		authorized.GET("/barang/:id", func(c *gin.Context) {
			id := c.Param("id")
			barang := barangService.GetBarang(id)
			if barang == nil {
				c.JSON(http.StatusNotFound, gin.H{"message": "Barang not found"})
				return
			}
			c.JSON(http.StatusOK, barang)
		})

		authorized.PUT("/barang/:id", func(c *gin.Context) {
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

		authorized.DELETE("/barang/:id", func(c *gin.Context) {
			id := c.Param("id")
			deleted := barangService.DeleteBarang(id)
			if !deleted {
				c.JSON(http.StatusNotFound, gin.H{"message": "Barang not found"})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Barang deleted"})
		})
	}

	r.POST("/register", userService.Register)
	r.POST("/login", userService.Login)
	r.Run(":" + APP_PORT)
}

package handlers

import (
	"bioskop-app/database"
	"bioskop-app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateBioskop(c *gin.Context) {
	var bioskop models.Bioskop
	if err := c.ShouldBindJSON(&bioskop); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nama dan Lokasi wajib diisi"})
		return
	}
	result := database.DB.Create(&bioskop)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Bioskop berhasil ditambahkan", "data": bioskop})
}

package productcontroller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tentangkode/go-restapi-gin/models"
	"gorm.io/gorm"
)

func Index(c *gin.Context) {

	var products []models.Produk

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"Produk": products})
}

func Show(c *gin.Context) {

	var products models.Produk
	id := c.Param("id")

	if err := models.DB.First(&products, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data gak adaa"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"Produk:": products})
}

func Create(c *gin.Context) {

	var products models.Produk

	if err := c.ShouldBindJSON(&products); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	models.DB.Create(&products)
	c.JSON(http.StatusOK, gin.H{"produk": products})
}

func Update(c *gin.Context) {

	var products models.Produk
	id := c.Param("id")

	if err := c.ShouldBindJSON(&products); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&products).Where("id=?", id).Updates(&products).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak bisa update"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "data berhasil di ubah"})
}

func Delete(c *gin.Context) {

	var products models.Produk

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&products); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if models.DB.Delete(&products, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak bisa hapus"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "data berhasil di hapus"})
}

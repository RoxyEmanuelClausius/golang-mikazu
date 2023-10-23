package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tentangkode/go-restapi-gin/controllers/productcontroller"
	"github.com/tentangkode/go-restapi-gin/models"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/api/produk", productcontroller.Index)
	r.GET("/api/produk/:id", productcontroller.Show)
	r.POST("/api/produk", productcontroller.Create)
	r.PUT("/api/produk", productcontroller.Update)
	r.DELETE("/api/produk", productcontroller.Delete)

	r.Run()
}

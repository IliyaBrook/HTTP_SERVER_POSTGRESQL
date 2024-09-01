package routes

import (
	"github.com/gin-gonic/gin"
	"main/internal/handlers/products"
)

func RegisterProductsRoutes(r *gin.Engine) {
	route := WithMwsRoute(r)
	{
		route.POST("/products", products.AddProduct)
		route.DELETE("/products", products.DeleteProduct)
		route.PATCH("/products", products.UpdateProduct)
		route.GET("/products", products.GetProducts)
	}
}

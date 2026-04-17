package routes

import (
	"go-pagination/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	productGroup := r.Group("/products")
	{
		productGroup.GET("/offset", controllers.GetOffsetPagination)
		productGroup.GET("/cursor", controllers.GetCursorPagination)
		productGroup.GET("/search", controllers.GetDynamicSearch)
	}
}

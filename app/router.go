package app

import (
	"submission-2/controller"

	"github.com/gin-gonic/gin"
)

func OrderRoute(orderController controller.OrderController) *gin.Engine {
	router := gin.Default()

	order := router.Group("/orders")
	order.GET("", orderController.FindAll)
	order.POST("", orderController.Create)
	order.GET("/:id", orderController.FindById)
	order.PUT("/:id", orderController.Update)
	order.DELETE("/:id", orderController.Delete)

	return router
}

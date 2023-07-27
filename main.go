package main

import (
	"submission-2/app"
	"submission-2/controller"
	"submission-2/repository"
	"submission-2/service"
)

func main() {
	db := app.NewDB()
	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository)
	orderController := controller.NewOrderController(orderService)

	router := app.OrderRoute(orderController)

	router.Run()
}

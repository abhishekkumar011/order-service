//This file is responsible for registering all the API routes. (like index.js in express.js where you define endpoints)

package routes

import (
	"food-delivery-order/internal/handlers"
	"food-delivery-order/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {

	orderService := &service.OrderService{}

	orderHandler := &handlers.OrderHandler{
		Service: orderService,
	}

	router.POST("/events", orderHandler.HandleEvent)

	router.GET("/orders", orderHandler.GetOrders)
}

/*
	1. What is this file?
	   - Create the service
	   - Create the handler
	   - Connect URLs to handler methods

	2. func RegisterRoutes(router *gin.Engine)
	   - This function receives the Gin router. *gin.Engine is the main router of your application.

	3. orderService := &Service.OrderService{} - Creates an instance(object) of OrderService.
	   - & - Create an OrderService and return its pointer.

	4. orderHandler := &handlers.OrderHandler{
          Service: orderService,
       }
		  - Creates an OrderHandler while creating it we pass the service
		  - We are giving the handler the service it needs.
*/

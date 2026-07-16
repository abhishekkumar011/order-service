//The handler is the part of your application that receives HTTP requests and sends HTTP responses.
//A handler is the entry point for an HTTP request. It receives the request, converts the JSON into Go structs, calls the service to perform the business logic, and sends the appropriate HTTP response back to the client.

package handlers

import (
	"net/http"

	"food-delivery-order/internal/models"
	"food-delivery-order/internal/service"

	"github.com/gin-gonic/gin"
)

// The handler stores a reference to the service. it stores the OrderService
type OrderHandler struct {
	Service *service.OrderService
}

// This is a method HandleEvent that belongs to OrderHandler. It receives the current HTTP request/response (gin.Context) so it can process the request and send a response.
func (h *OrderHandler) HandleEvent(c *gin.Context) {

	var event models.Event

	if err := c.ShouldBindJSON(&event); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := h.Service.ProcessEvent(event)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Event processed successfully",
	})
}

func (h *OrderHandler) GetOrders(c *gin.Context) {

	orders, err := h.Service.GetOrders()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, orders)
}

/*
	1. func (h *OrderHandler) HandleEvent(c *gin.Context)
	   - h = the handler instance (just a variable name) think of like this in js (this.service → h.Service)
	   - c = the HTTP request and response(* gin.context)
	   - async function handleEvent(req, res) {} - equivalent to Express
	   - (h *OrderHandler) This is called the receiver. means This function belongs to OrderHandler.
	     - *OrderHandler means use the Originl OrderHandler not copy

	2. var event models.Event - Creates an empty Event (Object).
	3. if err := c.ShouldBindJSON(&event); err != nil
	   - ShouldBindJSON() reads the incoming JSON and fills the event struct (Covert it into your Go struct).
	   - &event - It is the address of event
	     - means Here is the object. Please write the JSON data into it.
		 - In Go, structs are copied by default, so you pass a pointer (&event) to let the function modify the original struct.
	   - Why the semicolon (;) - The first part runs once, then the condition is checked.

	4. err := h.Service.ProcessEvent(event)
	   - Ask the OrderService to process this event, and store any error it returns in the variable err.
*/

//This service receives an order event, checks what kind of event it is (create, update status, or update items), and performs the corresponding MongoDB operation.

package service

import (
	"context"
	"time"

	"food-delivery-order/internal/database"
	"food-delivery-order/internal/models"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

type OrderService struct{}

func (s *OrderService) ProcessEvent(event models.Event) error {

	switch event.Type {

	case "order.create":
		return s.createOrder(event)

	case "order.update.status":
		return s.updateStatus(event)

	case "order.update.items":
		return s.updateItems(event)

	}

	return nil
}

func (s *OrderService) createOrder(event models.Event) error {
	order := models.Order{
		OrderID:      uuid.New().String(),
		CustomerID:   event.CustomerID,
		RestaurantID: event.RestaurantID,
		Status:       "Pending",
		Items:        event.Items,
		LastUpdated:  time.Now(),
	}

	_, err := database.OrderCollection.InsertOne(
		context.Background(),
		order,
	)

	return err
}

func (s *OrderService) updateStatus(event models.Event) error {

	_, err := database.OrderCollection.UpdateOne(
		context.Background(),

		bson.M{
			"orderId": event.OrderID,
		},

		bson.M{
			"$set": bson.M{
				"status":      event.Status,
				"lastUpdated": time.Now(),
			},
		},
	)

	return err
}

func (s *OrderService) updateItems(event models.Event) error {

	_, err := database.OrderCollection.UpdateOne(
		context.Background(),

		bson.M{
			"orderId": event.OrderID,
		},

		bson.M{
			"$set": bson.M{
				"items":       event.Items,
				"lastUpdated": time.Now(),
			},
		},
	)

	return err
}

/*
	1. type OrderService struct{} Think of it like a class in TypeScript.
	2. bson.M is just a MongoDB map. It's like a JavaScript object: It's used to build MongoDB queries and updates.
	3. in the func (s *OrderService) updateItems(event models.Event)
	   - (s *OrderService) This is called the receiver. 
	     - s is just a variable name, like this in JavaScript.
		 - The * means pointer. menas Use the original OrderService, not a copy.
	   - updateItems is a method that belongs to OrderService. It's similar to a class method in JavaScript/TypeScript.
	   - (event models.event) This is the parameter.
	     - variable name: event
		 - Type: models.Event
	   - error means the return type of function but if everything is succed then it returns nil means no error	 
*/

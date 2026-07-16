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

func (s *OrderService) GetOrders() ([]models.Order, error) {

	cursor, err := database.OrderCollection.Find(
		context.Background(),
		bson.M{},
	)

	if err != nil {
		return nil, err
	}

	var orders []models.Order

	if err := cursor.All(context.Background(), &orders); err != nil {
		return nil, err
	}

	return orders, nil
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

	4. ([]models.Order, error) - Returns two Values 1. list of orders 2. error
	    - bson.M{} inside M we can define filter based on that we can fetch documents
		  - but {} means no filter it tells MongoDB Give me every document in this collection.  
	
	5. cursor, err := database.OrderCollection.Find(
		context.Background(),
		bson.M{},
	)
		- Think of a cursor as a bookmark or iterator.
		- When you run this func you will get cursor instead of orders[] 
		- Find() in go does not return the actual orders it returns a cursor.
		  - cursor - A pointer that lets you read the matching documents.
		  	- Go gives you something like this:
			- Cursor
               ↓
              Currently pointing at Page 1
			  - The cursor knows where the results are, but it hasn't copied them into Go memory yet.

		- Cusort.All() - Because now we want to read everything the cursor found. Convert cursor into an array of orders	  
		               - This means: Cursor read all the documents you're pointing to and copy them into this slice(order array).
*/

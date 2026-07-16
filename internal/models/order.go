package models

import "time"

type Order struct {
	OrderID      string    `json:"orderId" bson:"orderId"`
	CustomerID   string    `json:"customerId" bson:"customerId"`
	RestaurantID string    `json:"restaurantId" bson:"restaurantId"`
	Status       string    `json:"status" bson:"status"`
	Items        []Item    `json:"items" bson:"items"`
	LastUpdated  time.Time `json:"lastUpdated" bson:"lastUpdated"`
}


/*
	1. time.Time is like Date in JavaScript
*/
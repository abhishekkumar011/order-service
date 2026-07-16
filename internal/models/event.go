package models

type Event struct {
	Type         string `json:"type"`
	CustomerID   string `json:"customerId,omitempty"`
	RestaurantID string `json:"restaurantId,omitempty"`
	OrderID      string `json:"orderId,omitempty"`
	Status       string `json:"status,omitempty"`
	Items        []Item `json:"items,omitempty"`
}


/*
	1. omitempty - `json:"status,omitempty"`
	             - If it's empty, don't include it in the JSON.
*/
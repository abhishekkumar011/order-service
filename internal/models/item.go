package models

type Item struct {
	ItemID string `json:"itemId" bson:"itemId"`
	Qty    int    `json:"qty" bson:"qty"`
}


/*
	1. "json:"itemId" bson:"itemId"" These are called struct tags.
	   - Go uses capitalization to control visibility. like ItemID
	   - API conventions excepts the name in camelCase that's why we use struct tags
	   - json: "itemId" Used when sending or receiving JSON through your API
	   - bson: "itemId" Used when storing or reading data from MongoDB
*/
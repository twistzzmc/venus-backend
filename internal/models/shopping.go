package models

import (
	"time"
)

type ShoppingList struct {
	Name  string             `bson:"name"`
	Items []ShoppingListItem `bson:"items"`
	Owner User               `bson:"owner"`
}

type ShoppingListItem struct {
	Name     string     `bson:"name"`
	IsBought bool       `bson:"is_bought"`
	BuyTime  *time.Time `bson:"buy_time"`
}

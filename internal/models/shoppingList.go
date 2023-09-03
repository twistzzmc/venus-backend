package models

type ShoppingListItem struct {
	Name      string `bson:"name"`
	IsBought  bool   `bson:"is_bought"`
	IsDeleted bool   `bson:"is_deleted"`
}

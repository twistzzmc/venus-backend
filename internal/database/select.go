package database

import (
	"context"
	"venus/internal/models"

	"go.mongodb.org/mongo-driver/bson"
)

func FindAllShoppingLists() (lists []models.ShoppingList, err error) {
	coll := getShoppingListColl()
	cursor, err := coll.Find(context.TODO(), bson.D{})
	if err != nil {
		return
	}

	err = cursor.All(context.TODO(), &lists)
	return
}

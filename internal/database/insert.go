package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	"venus/internal/models"
)

func InsertShoppingListItem(client *mongo.Client, item models.ShoppingListItem) error {
	coll := client.Database("venus").Collection("shopping_list")

	res, err := coll.InsertOne(context.TODO(), item)
	if err != nil {
		return err
	}

	fmt.Printf("inserted document with ID %v\n", res.InsertedID)
	return nil
}

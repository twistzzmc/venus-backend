package database

import "go.mongodb.org/mongo-driver/mongo"

func getShoppingListColl() *mongo.Collection {
	return client.Database("venus").Collection("shopping_list")
}

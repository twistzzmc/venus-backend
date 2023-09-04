package main

import (
	"fmt"

	"net/http"
	"venus/internal/config"
	"venus/internal/database"

	"github.com/gin-gonic/gin"
)

func ptr[T any](x T) *T {
	return &x
}

func main() {
	dbConf := config.ReadConfigFile()
	database.ConnectToMongoDB(dbConf)
	defer database.DisconnectFromMongoDB()

	fmt.Println(dbConf)
	fmt.Println("Pomyślnie połączono do bazy danych")

	router := gin.Default()
	router.GET("/shoppingLists", getShoppingLists)

	router.Run("localhost:8080")
}

func getShoppingLists(c *gin.Context) {
	shoppingLists, err := database.FindAllShoppingLists()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Nie udało się pobrać list zakupów")
		return
	}

	c.IndentedJSON(http.StatusOK, shoppingLists)
}

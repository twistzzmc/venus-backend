package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"venus/internal/customerrors"
	"venus/internal/logger"
	"venus/internal/models"

	"venus/internal/config"
	"venus/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	logger.Logger = logger.ConsoleLogger{}

	configuration := config.ReadConfig()
	database.TestConnection(configuration.Database)

	fmt.Println(configuration.Database)
	fmt.Println("Pomyślnie połączono do bazy danych")

	router := gin.Default()
	router.Use(RecoveryMiddleware)
	router.Use(ginBodyLogMiddleware)
	//router.GET("/shoppingLists", getShoppingLists)
	router.PUT("/register", register)

	if err := router.Run("localhost:8080"); err != nil {
		panic(err)
	}
}

func register(context *gin.Context) {
	var registerUser models.RegisterUser

	if err := context.ShouldBindJSON(&registerUser); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			jsonError := customerrors.JSONError{
				Field: ve[0].Field(),
				Tag:   ve[0].Tag(),
			}

			fmt.Println(jsonError)
			context.JSON(http.StatusBadRequest, jsonError.ToJSON())
			return
		}

		panic(err)
		return
	}

	fmt.Println(registerUser)
	context.IndentedJSON(http.StatusCreated, registerUser)
}

func RecoveryMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if panicErr := recover(); panicErr != nil {
				logger.Logger.LogPanic(panicErr)
				context.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		context.Next()
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func ginBodyLogMiddleware(c *gin.Context) {
	blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
	c.Writer = blw
	c.Next()
	statusCode := c.Writer.Status()
	if statusCode >= 400 {
		//ok this is an request with error, let's make a record for it
		// now print body (or log in your preferred way)
		fmt.Println("Response body: " + blw.body.String())
	}
}

//func getShoppingLists(c *gin.Context) {
//	shoppingLists, err := database.FindAllShoppingLists()
//	if err != nil {
//		c.IndentedJSON(http.StatusInternalServerError, "Nie udało się pobrać list zakupów")
//		return
//	}
//
//	c.IndentedJSON(http.StatusOK, shoppingLists)
//}

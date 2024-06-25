package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"pizzas.com/api/models"
)

func main() {
	server := gin.Default()

	server.GET("/pizzas", HomeController)
	server.POST("/pizzas", createPizza)

	server.Run(":8080")
}

func HomeController(context *gin.Context) {
	pizzas := models.GetAllPizzas()
	context.JSON(http.StatusOK, pizzas)
}

func createPizza(context *gin.Context) {
	var pizza models.Pizza
	err := context.ShouldBindJSON(&pizza)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	pizza.ID = 1
	pizza.CreateOn = time.Now()

	pizza.Save()

	context.JSON(http.StatusCreated, pizza)
}

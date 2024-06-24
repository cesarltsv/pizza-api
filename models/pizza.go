package models

import "time"

type Pizza struct {
	ID          int
	Name        string   `binding:"required"`
	Description string   `binding:"required"`
	Ingredients []string `binding:"required"`
	CreateOn    time.Time
	price       float64 `binding:"required"`
}

var pizzas = []Pizza{}

func (p Pizza) Save() {
	// TODO: add logic
	pizzas = append(pizzas, p)
}

func GetAllPizzas() []Pizza {
	return pizzas
}

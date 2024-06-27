package models

import (
	"time"

	"pizzas.com/api/db"
)

type Pizza struct {
	ID          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Ingredients string `binding:"required"`
	CreateOn    time.Time
	price       float64 `binding:"required"`
}

var pizzas = []Pizza{}

func (p Pizza) Save() error {
	query := `
		INSERT INTO pizzas(name, description, ingredients, createOn, price)
		VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(p.Name, p.Description, p.Ingredients, p.CreateOn, p.price)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	p.ID = id
	return err
}

func GetAllPizzas() []Pizza {
	return pizzas
}

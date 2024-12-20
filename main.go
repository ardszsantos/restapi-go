package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Pizza struct {
	ID    int
	nome  string
	preco float64
}

func main() {

	router := gin.Default()
	router.GET("/pizzas", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"pizzas": "pizzas, pizas, pisas",
		})
	})

	var pizzas = []Pizza{
		{ID: 1, nome: "Toscana", preco: 49.5},
		{ID: 2, nome: "Marguerita", preco: 79.5},
		{ID: 3, nome: "Atum com queijo", preco: 69.5},
	}

	fmt.Println(pizzas)
	router.Run()
}

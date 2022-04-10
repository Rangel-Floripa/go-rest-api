package main

import (
	"api-go-rest2/models"
	"api-go-rest2/routes"
	"fmt"
)

func main() {
	models.Personalidades =[]models.Personalidade{
		{Id: 1, Nome: "Nome1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome2", Historia: "Historia 2"},

	}
	fmt.Println("Iniciando o gervidor Rest com Go ...")
	routes.HandleResquest()

}

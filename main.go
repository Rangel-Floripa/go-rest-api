package main

import (
	"api-go-rest2/routes"
	"fmt"
)

func main() {
	fmt.Println("Iniciando o gervidor rest com Go")
	routes.HandleResquest()

}

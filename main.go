package main

import (
	"fmt"

	"github.com/RodrigoMatsumoto/go-rest-api/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}

package main

import (
	"fmt"
	"programming-lang-api/models"
	"programming-lang-api/routes"
)

// main: Iniciando API chama o pacote routes.HandleRequest() que contem as instruções a serem seguidas em cada rota da API
func main() {
	fmt.Println("Iniciando api...")

	models.ProgrammingLanguages = []models.ProgrammingLangs{
		{Id: 1, Name: "Python", Typed: false, History: "python..."},
		{Id: 2, Name: "Golang", Typed: false, History: "golang..."},
	}

	fmt.Println(models.ProgrammingLanguages[0].Name)

	routes.HandleRequest()

}

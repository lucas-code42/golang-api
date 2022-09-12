package routes

import (
	"log"
	"net/http"
	"programming-lang-api/controllers"
)

// HandleRequest abre uma porta para a execução da API e faz o mapeamento do que cada rota é resposável
func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/languages", controllers.ReadLanguages)
	http.HandleFunc("/create/languages", controllers.CreateLanguages)
	http.HandleFunc("/update", controllers.UpdateLanguages)

	// abre a porta 5000 para o servidor rodar
	// precisa ser a última linha
	log.Fatal(http.ListenAndServe(":5000", nil))
}

package controllers

/*
CRIAR UMA INCREMENTADOR DE ID PARA STRUCT
CRIAR UMA FUNÇÃO SANITÁRIA QUE VAI TRATAR OS INPUTS DO USUARIO DENTRO DA API
FUNÇÃO UPDATE: Criar uma verificação caso o id passado nao exista
MIDLEWARE PARA STATUS CODE
MIDDLEWARE PARA SET HEADERS CONTENT TYPE
CRIAR UMA FUNCAO PARA LOGS
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"programming-lang-api/models"
	"strconv"
)

// Home executa o processamento que veio encaminhado da rota "/"
func Home(w http.ResponseWriter, r *http.Request) {
	// seta no headers que a resposta é em json
	w.Header().Set("content-type", "application/json")

	// validação do método vindo do client
	if r.Method == "GET" {
		fmt.Fprint(w, "home page")
	} else {
		fmt.Println(r.Method)
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// Mostra uma lista de lingaguems cadastradas "/languages"
func ReadLanguages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	// verifica o método usado na requisição, caso nao seja retorna method not allowed
	if r.Method != "GET" {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
	fmt.Println("Processamento ShowLanguages")

	// Retorna o objeto instanciado em models em JSON
	json.NewEncoder(w).Encode(models.ProgrammingLanguages)
}

// CreateLanguages: cria novas instancias da struct models.ProgrammingLangs
func CreateLanguages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "invalid method", http.StatusMethodNotAllowed)
		return
	}

	var newLang models.ProgrammingLangs

	// maneira como converter um json em um array de bytes, body representa um array de bytes e err um erro.
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid body", 400)
	}

	// aqui será guardado um valor booleano para saber se o body enviado é true ou false
	var bodyIsValid bool

	// fazendo Unmarshal, ou seja pegando um array de bytes (decode) e alocando o seu valor na referência de &newLang e verificando o erro
	JSONerr := json.Unmarshal(body, &newLang)
	if JSONerr != nil {
		http.Error(w, "Invalid body", 400)
		bodyIsValid = false
	} else {
		// body valido entao seguimos com execução
		bodyIsValid = true
	}

	// se o body for válido adciona um objeto dentro da instancia models.ProgrammingLanguages
	if bodyIsValid {
		models.ProgrammingLanguages = append(models.ProgrammingLanguages, newLang)
		// Retornando um http status code
		w.WriteHeader(http.StatusOK)
		// Retornando uma msg para o client
		w.Write([]byte("Sucess!"))
	}

	// r.Body http cru
	// body []bytes
	// newLang json após Unmarshal
}

// UpdateLanguages faz o update de uma determinada linguagem com base no ID
func UpdateLanguages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	if r.Method != "POST" {
		http.Error(w, "Invalid update", http.StatusMethodNotAllowed)
		return
	}
	// capturando uma query string na url, por padrão ela vem como string
	id := r.URL.Query().Get("id")
	history := r.URL.Query().Get("history")

	// parse de string para int
	idInt, errInt := strconv.Atoi(id)
	if errInt != nil {
		http.Error(w, "invalid query string", 400)
		return
	}

	// iteração sobre a instancia da struct, caso o ID contenha dentro da struct então a modificacao no campo history é permitida
	var indexTochange int
	for index, value := range models.ProgrammingLanguages {
		if value.Id == idInt {
			indexTochange = index

		}
	}

	// acessando o indice e depois a chave dentro de um struct
	models.ProgrammingLanguages[indexTochange].History = history
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Sucess!"))
}

// deleta um elemento da lista de lingaguens de programação
func DeleteLanguage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	if r.Method != "DELETE" {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("id")
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		http.Error(w, "invalid query string", 400)
		return
	}

	var removeIndex int
	for index, value := range models.ProgrammingLanguages {
		if value.Id == idInt {
			removeIndex = index
		}
	}
	models.ProgrammingLanguages = removeLangByIndex(models.ProgrammingLanguages, removeIndex)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Sucess!"))
}

// removeLangByIndex delta um elemento da lista atraves de um indice
func removeLangByIndex(s []models.ProgrammingLangs, index int) []models.ProgrammingLangs {
	return append(s[:index], s[index+1:]...)
}

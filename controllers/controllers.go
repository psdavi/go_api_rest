package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-rest-api/database"
	"go-rest-api/models"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

// Buscando do banco -------------------------------------------------------------------------
// FIND ALL
func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

// FIND ONE
// first pega o primeiro resultado, ou seja o id, e como ele é único, sempre vai vir o esperado
func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var p models.Personalidade
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

// CREATE
func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novaPersonalidade)
	database.DB.Create(&novaPersonalidade)
	json.NewEncoder(w).Encode(novaPersonalidade)
}

// DELETE
func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

// EDIT
func EditaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}

// Buscando do banco --------------------------------------------------

// Buscando informações mockadas -----------------------------------------------------------

//func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
//	json.NewEncoder(w).Encode(models.Personalidades)
//}

//func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
//	vars := mux.Vars(r)
//	id := vars["id"]
//
//	for _, personalidade := range models.Personalidades {
//		if strconv.Itoa(personalidade.Id) == id {
//			json.NewEncoder(w).Encode(personalidade)
//		}
//	}
//}
// Buscando informações mockadas -----------------------------------------------------------

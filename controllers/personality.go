package controllers

import (
	"alura-go-personalities/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func FindAll(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.FindAllPersonalities())
}

func FindById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return
	}
	personality := models.FindPersonalityById(id)
	if personality.Id == 0 {
		return
	}
	json.NewEncoder(w).Encode(personality)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	json.NewDecoder(r.Body).Decode(&personality)
	models.CreatePersonality(&personality)
	json.NewEncoder(w).Encode(personality)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return
	}
	p := models.DeletePersonality(id)
	json.NewEncoder(w).Encode(p)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return
	}
	var personality models.Personality
	json.NewDecoder(r.Body).Decode(&personality)
	p := models.UpdatePersonality(id, personality)
	json.NewEncoder(w).Encode(p)
}

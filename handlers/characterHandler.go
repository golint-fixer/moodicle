package handlers

import (
	"encoding/json"
	"github.com/mattdotmatt/moodicle/models"
	"github.com/mattdotmatt/moodicle/repositories"
	"gopkg.in/validator.v2"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

/*
	Get all the characters in the database
*/
func GetCharacters(characters repositories.CharacterRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		p := mux.Vars(r)

		tester := p["name"]

		log.Println(tester)

		c, err := characters.Character("1")

		if characters == nil || err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(c)
	}
}

/*
	Save a payload of characters to the database. These replace the existing items in the store
*/
func SaveCharacters(characters repositories.CharacterRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		p := mux.Vars(r)
		tester := p["name"]

		log.Println(tester)

		decoder := json.NewDecoder(r.Body)

		var input models.Character

		err := decoder.Decode(&input)

		// Validate input
		if err := validator.Validate(input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode("First name cannot be empty")
			return
		}

		input.Id = "3"

		if err = characters.SaveCharacter(input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}

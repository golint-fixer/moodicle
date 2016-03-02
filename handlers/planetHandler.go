package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mattdotmatt/moodicle/models"
	"github.com/mattdotmatt/moodicle/repositories"
	"github.com/satori/go.uuid"
	"gopkg.in/validator.v2"
	"net/http"
)

/*
	Get a planet
*/
func GetPlanet(planets repositories.PlanetRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !authenticate(r) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		p := mux.Vars(r)

		owner := p["name"]
		id := p["id"]

		c, err := planets.GetPlanet(owner, id)

		if planets == nil || err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(c)
	}
}

/*
	Get all planets
*/
func GetPlanets(planets repositories.PlanetRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !authenticate(r) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		p := mux.Vars(r)

		owner := p["name"]

		c, err := planets.GetPlanets(owner)

		if planets == nil || err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(c)
	}
}

/*
	Save a planet
*/
func SavePlanet(planets repositories.PlanetRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !authenticate(r) {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		p := mux.Vars(r)

		owner := p["name"]

		decoder := json.NewDecoder(r.Body)

		var input models.Planet

		err := decoder.Decode(&input)

		// Validate input
		if err := validator.Validate(input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}

		input.Id = uuid.NewV1().String()

		if err = planets.SavePlanet(owner, input); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		json.NewEncoder(w).Encode(input.Id)
		w.WriteHeader(http.StatusOK)
	}
}

func authenticate(r *http.Request) bool {

	apiKey := r.Header.Get("API_KEY")

	if apiKey != "1234" {
		return false
	}

	return true
}

package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/mattdotmatt/moodicle/models"
	"io/ioutil"
	"os"
)

type PlanetRepository interface {
	GetPlanets(owner string) ([]models.Planet, error)
	GetPlanet(owner string, id string) (models.Planet, error)
	SavePlanet(owner string, planet models.Planet) error
}

type planetRepository struct {
	FileLocation string
}

func NewPlanetRepository(fileLocation string) PlanetRepository {
	repo := planetRepository{FileLocation: fileLocation}
	return &repo
}

func (db *planetRepository) GetPlanets(owner string) ([]models.Planet, error) {

	planets := []models.Planet{}

	folder := fmt.Sprintf("%s/%s/", db.FileLocation, owner)

	files, _ := ioutil.ReadDir(folder)

	var err error

	for _, f := range files {
		planet := models.Planet{}

		file := fmt.Sprintf("%s/%s", folder, f.Name())

		content, err := ioutil.ReadFile(file)
		if err == nil {
			err = json.Unmarshal(content, &planet)

			if err == nil {
				planets = append(planets, planet)
			}
		}
	}

	return planets, err
}

func (db *planetRepository) GetPlanet(owner string, id string) (models.Planet, error) {

	planet := models.Planet{}

	folder := fmt.Sprintf("%s/%s", db.FileLocation, owner)

	file := fmt.Sprintf("%s/%s.json", folder, id)

	content, err := ioutil.ReadFile(file)

	if err == nil {
		err = json.Unmarshal(content, &planet)
	}

	return planet, err
}

func (db *planetRepository) SavePlanet(owner string, planet models.Planet) error {

	folder := fmt.Sprintf("%s/%s", db.FileLocation, owner)

	err := checkFolderExists(folder)

	c, err := json.Marshal(planet)

	if err == nil {
		file := fmt.Sprintf("%s/%s.json", folder, planet.Id)
		err = ioutil.WriteFile(file, c, 0644)
	}

	return err
}

func checkFolderExists(fileLocation string) error {
	if _, err := ioutil.ReadDir(fileLocation); err != nil {
		if err = os.Mkdir(fileLocation, 1); err != nil {
			return err
		}
	}
	return nil
}

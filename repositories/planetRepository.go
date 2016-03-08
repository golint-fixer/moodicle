package repositories

import (
	"encoding/json"
	"fmt"
	"github.com/blang/vfs"
	"github.com/mattdotmatt/moodicle/models"
	"os"
)

type PlanetRepository interface {
	GetPlanets(owner string) ([]models.Planet, error)
	GetPlanet(owner string, id string) (models.Planet, error)
	SavePlanet(owner string, planet models.Planet) error
}

type planetRepository struct {
	vfs.Filesystem
}

var f string

func NewPlanetRepository(folder string, fileSystem vfs.Filesystem) PlanetRepository {
	f = folder
	return &planetRepository{fileSystem}
}

func (db *planetRepository) GetPlanets(owner string) ([]models.Planet, error) {

	planets := []models.Planet{}

	folder := fmt.Sprintf("%s/%s/", f, owner)

	files, _ := db.ReadDir(folder)

	var err error

	for _, f := range files {

		planet, err := readPlanet(db, fmt.Sprintf("%s%s", folder, f.Name()))

		if err == nil {
			planets = append(planets, planet)
		}
	}

	return planets, err
}

func (db *planetRepository) GetPlanet(owner string, id string) (models.Planet, error) {

	return readPlanet(db, fmt.Sprintf("%s/%s/%s.json", f, owner, id))

}

func (db *planetRepository) SavePlanet(owner string, planet models.Planet) error {

	folder := fmt.Sprintf("%s/%s", f, owner)

	err := checkFolderExists(db, folder)

	c, err := json.Marshal(planet)

	if err == nil {
		fs, err := db.OpenFile(fmt.Sprintf("%s/%s.json", folder, planet.Id), os.O_CREATE|os.O_RDWR, 0666)

		if err == nil {
			_, err = fs.Write([]byte(c))
		}
	}

	return err
}

func checkFolderExists(db *planetRepository, fileLocation string) error {
	if _, err := db.ReadDir(fileLocation); err != nil {
		if err = db.Mkdir(fileLocation, 1); err != nil {
			return err
		}
	}
	return nil
}

func readPlanet(db *planetRepository, file string) (models.Planet, error) {

	var planet models.Planet

	content, err := db.OpenFile(file, os.O_RDONLY, 0)

	if err != nil {
		return planet, err
	}

	stats, _ := db.Stat(file)

	p := make([]byte, stats.Size())

	_, err = content.Read(p)

	if err != nil {
		return planet, err
	}

	json.Unmarshal(p, &planet)

	return planet, err
}

package repositories

import (
	"github.com/mattdotmatt/moodicle/models"
	"github.com/boltdb/bolt"
	"encoding/gob"
	"bytes"
)

type CharacterRepository interface {
	Character(id string) (models.Character, error)
	SaveCharacter(character models.Character) error
}

type characterRepository struct {
	*bolt.DB `inject:""`
}

func NewCharacterRepository() CharacterRepository {
	return &characterRepository{}
}

func (db *characterRepository) Character(id string) (models.Character, error) {
	var value_variable []byte

	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("posts"))
		value_variable = b.Get([]byte(id))

		return nil
	})


	char := models.Character{}
	pCache := bytes.NewBuffer(value_variable)
	dec := gob.NewDecoder(pCache) // Will read from network.
	dec.Decode(char)
	return char, err
}

func (db *characterRepository) SaveCharacter(character models.Character) error {

	char := []byte{}
	pCache := bytes.NewBuffer(char)
	dec := gob.NewEncoder(pCache) // Will read from network.
	dec.Encode(character)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("posts"))
		v := b.Put([]byte(character.Id), []byte("test"))
		return v
	})

	return nil
}

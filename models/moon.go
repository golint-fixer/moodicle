package models

type Moon struct {
	Name string `json:"name" validate:"nonzero"`
}

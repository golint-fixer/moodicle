package models

type Character struct {
	Id    string
	FirstName string `json:"firstName" validate:"nonzero"`
	LastName  string `json:"lastName"`
}

package models

type Planet struct {
	Id         string `json:"id"`
	Name       string `json:"name,omitempty" validate:"nonzero"`
	HoursInDay int    `json:"hoursInDay,omitempty" validate:"min=1,max=999999"`
	Moons      []Moon `json:"moons,omitempty"`
}

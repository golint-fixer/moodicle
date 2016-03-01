package models

type Planet struct {
	Id         string `json:"id"`
	Name       string `json:"name" validate:"nonzero"`
	HoursInDay int    `json:"hoursInDay" validate:"min=1,max=999999"`
	Moons      []Moon `json:"moons"`
}

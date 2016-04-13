package repositories

import (
	"github.com/mattdotmatt/moodicle/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetCharacters(t *testing.T) {

	Convey("Given planets exist", t, func() {

		repository := planetRepository{dummyFS{}}

		Convey("When I request all existing planet", func() {

			planets, err := repository.GetPlanets("bob")

			Convey("Then I get back the requested planets", func() {

				So(err, ShouldBeNil)
				So(len(planets), ShouldEqual, 2)
				So(planets[0].Id, ShouldEqual, "1234")
				So(planets[1].Id, ShouldEqual, "5678")
			})

		})

		Convey("When I request an existing planet", func() {

			planet, err := repository.GetPlanet("bob", "1234")

			Convey("Then I get back the requested planet", func() {

				So(err, ShouldBeNil)
				So(planet.Id, ShouldEqual, "1234")
			})

		})

		Convey("When I delete an existing planet", func() {

			planet, err := repository.GetPlanet("bob", "1234")

			So(err, ShouldBeNil)
			So(planet.Id, ShouldEqual, "1234")

			err = repository.DeletePlanet("bob", "1234")

			So(err, ShouldBeNil)

			planet, err = repository.GetPlanet("bob", "1234")

			Convey("Then I should not get back the requested planet", func() {

				So(err, ShouldBeNil)
				So(planet.Id, ShouldEqual, "1234")
			})

		})

		Convey("When I request an unknown planet", func() {

			planet, err := repository.GetPlanet("bob", "UNKNOWN")

			Convey("Then I get back nil", func() {

				So(err.Error(), ShouldEqual, "An error")
				So(planet, ShouldResemble, models.Planet{})
			})

		})
	})
}

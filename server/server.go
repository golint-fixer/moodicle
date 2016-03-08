package server

import (
	"fmt"
	"github.com/blang/vfs"
	"github.com/codegangsta/negroni"
	"github.com/facebookgo/inject"
	"github.com/mattdotmatt/moodicle/repositories"
	"github.com/mattdotmatt/moodicle/routers"
	"log"
	"net/http"
)

func Start(port int, folder string) {

	var router routers.Router
	var graph inject.Graph

	var osfs vfs.Filesystem = vfs.OS()

	// Setup DI
	if err := graph.Provide(
		&inject.Object{Value: repositories.NewPlanetRepository(folder, osfs)},
		&inject.Object{Value: &router}); err != nil {
		log.Fatalf("Error providing dependencies: ", err.Error())
	}

	if err := graph.Populate(); err != nil {
		log.Fatalf("Error populating dependencies: ", err.Error())
	}

	n := negroni.Classic()
	n.UseHandler(router.NewRouter())

	err := http.ListenAndServe(fmt.Sprintf(":%v", port), n)

	if err != nil {
		panic("Error: " + err.Error())
	}
}

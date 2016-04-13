package server

import (
	"github.com/blang/vfs"
	"github.com/codegangsta/negroni"
	"github.com/facebookgo/inject"
	"github.com/matryer/silk/runner"
	"github.com/mattdotmatt/moodicle/repositories"
	"github.com/mattdotmatt/moodicle/routers"
	"log"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestAPIEndpoint(t *testing.T) {
	// start a server
	var router routers.Router
	var graph inject.Graph

	var osfs vfs.Filesystem = vfs.OS()

	// Setup DI
	if err := graph.Provide(
		&inject.Object{Value: repositories.NewPlanetRepository("../server/data", osfs)},
		&inject.Object{Value: &router}); err != nil {
		log.Fatalf("Error providing dependencies: ", err.Error())
	}

	if err := graph.Populate(); err != nil {
		log.Fatalf("Error populating dependencies: ", err.Error())
	}

	n := negroni.Classic()
	n.UseHandler(router.NewRouter())

	s := httptest.NewServer(n)

	defer s.Close()

	// run all test files
	runner.New(t, s.URL).RunGlob(filepath.Glob("../specs/*.md"))
}

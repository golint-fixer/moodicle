package routers

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/mattdotmatt/moodicle/handlers"
	"github.com/mattdotmatt/moodicle/repositories"
	"net/http"
)

type Router struct {
	CharacterRepository repositories.CharacterRepository `inject:""`
}

func (router Router) NewRouter() *mux.Router {

	r := mux.NewRouter()

	r.PathPrefix("/api").Handler(negroni.New(
		negroni.HandlerFunc(ApiHeaderMiddleware),
		negroni.Wrap(apiRouter(router.CharacterRepository)),
	))

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/public/")))

	return r

}

func apiRouter(characters repositories.CharacterRepository) *mux.Router {

	r := mux.NewRouter()

	r.HandleFunc("/api/{name:[a-z]+}/characters", handlers.GetCharacters(characters)).Methods("GET")
	r.HandleFunc("/api/{name:[a-z]+}/characters", handlers.SaveCharacters(characters)).Methods("POST")

	return r
}

func ApiHeaderMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	w.Header().Set("Content-Type", "application/json")
	next(w, r)
}

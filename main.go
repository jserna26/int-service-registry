package main

import (
	"net/http"

	cfenv "github.com/cloudfoundry-community/go-cfenv"
	"github.com/gorilla/mux"
	"github.com/jserna26/SystemPOC/handlers/create"
	"github.com/jserna26/SystemPOC/handlers/getAll"
	"github.com/jserna26/SystemPOC/handlers/getOne"
	"github.com/jserna26/SystemPOC/repo"
	"github.com/rbrumby/scaffold"
)

func getDbUrl(servicename string, urikey string) (dburlStr string, err error) {
	env, err := cfenv.Current()

	if err != nil {
		panic(err)
	}
	//TODO Make DBService name a variable
	dbconf, err := env.Services.WithName(servicename)
	if err != nil {
		return
	}
	dburl := dbconf.Credentials[urikey]
	dburlStr = dburl.(string)
	return
}

func getInitRoutes(r repo.SystemRepo) (i scaffold.InitRoutes) {
	i = func(rt *mux.Router) {
		rt.Handle("/create", create.NewCreateSystemHandler(r)).Methods("POST")
		rt.Handle("/getAll", getAll.NewGetAllSystemsHandler(r)).Methods("GET")
		rt.Handle("/getOne", getOne.NewGetSystemHandler(r)).Methods("GET")
		rt.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(""))))
	}
	return
}

func main() {
	dburl, err := getDbUrl("localpsql", "uri")
	if err != nil {
		panic(err)
	}
	r, err := repo.NewPostgresSystemRepo(dburl)
	if err != nil {
		panic(err)
	}
	svr := scaffold.NewServer(getInitRoutes(r))
	svr.Run(":" + scaffold.GetPort())

}

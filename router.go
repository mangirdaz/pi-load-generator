package main

import (
	"net/http"

	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func NewRouter() {

	//create new router
	router := mux.NewRouter().StrictSlash(false)

	//api backend init
	router.Path("/healthz").Name("Health endpoint").HandlerFunc(http.HandlerFunc(Health))
	//api backend init
	apiV1 := router.PathPrefix("/api/v1").Subrouter()

	apiV1.Methods("GET").Path("/pi/{num}").Name("Pi Calc").Handler(http.HandlerFunc(Index))

	//middleware intercept
	midd := http.NewServeMux()
	midd.Handle("/", router)
	midd.Handle("/api/v1", negroni.New(
		negroni.Wrap(apiV1),
	))
	n := negroni.Classic()
	n.UseHandler(midd)
	url := fmt.Sprintf("%s:%s", "0.0.0.0", "8080")

	log.WithFields(log.Fields{
		"url": url,
	}).Debug("api: starting api server")

	log.Fatal(http.ListenAndServe(url, n))

	//return router

}

package main

import (
	"math"
	"net/http"
	"strconv"

	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
)

//Index index method for API
func Index(resp http.ResponseWriter, req *http.Request) {

	log.Debug("Index")
	vars := mux.Vars(req)
	log.Debugf("Calulating PI %s numbers", vars["num"])
	i, _ := strconv.Atoi(vars["num"])

	go pi(i)

}

//Health endpoit for app server
func Health(w http.ResponseWriter, r *http.Request) {
	log.Debug("/health endpoint called")
	w.Write([]byte("OK"))
}

//https://golang.org/doc/play/pi.go
//see: https://goo.gl/la6Kli
func pi(n int) float64 {
	ch := make(chan float64)
	for k := 0; k <= n; k++ {
		go term(ch, float64(k))
	}
	f := 0.0
	for k := 0; k <= n; k++ {
		f += <-ch
	}
	log.Debugf("Done. Pi %f", f)
	fmt.Print(f)
	return f
}

func term(ch chan float64, k float64) {
	ch <- 4 * math.Pow(-1, k) / (2*k + 1)
}

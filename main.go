package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/busStop", getBusStopsInfo).Methods("GET")
	router.HandleFunc("/busStop/urban/{id}", getUrbanBusStopInfo).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// getBusStopsInfo Gets all the information for busStops
func getBusStopsInfo(w http.ResponseWriter, r *http.Request) {
	// FIXME The whole function needs to be fixed
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://www.zaragoza.es/sede/servicio/urbanismo-infraestructuras/transporte-urbano/poste-autobus?rf=html&srsname=wgs84&start=0&rows=50&distance=500", nil)
	request.Header.Add("Accept", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		buildResultFail(w)
		return
	}
	defer resp.Body.Close()
	a, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", a)

	// Write results
	buildResultOK(w, a)
}

// getUrbanBusStopInfo Gets all the information for a specific urban bus stop
func getUrbanBusStopInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://www.zaragoza.es/sede/servicio/urbanismo-infraestructuras/transporte-urbano/poste-autobus/tuzsa-"+params["id"]+"?rf=html&srsname=wgs84", nil)
	request.Header.Add("Accept", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		buildResultFail(w)
		return
	}
	defer resp.Body.Close()
	a, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", a)
	// Write results
	buildResultOK(w, a)
}

func buildResultOK(w http.ResponseWriter, a []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(a)
}

func buildResultFail(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
}

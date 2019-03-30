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

	router.HandleFunc("/busStops", getBusStopsInfo).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

// getBusStopsInfo Gets all the information for a specific bus stop
func getBusStopsInfo(responseWriter http.ResponseWriter, r *http.Request) {
	//responseWriter.Header().Set("Accept", "application/json")
	client := &http.Client{}
	request, err := http.NewRequest("GET", "https://www.zaragoza.es/sede/servicio/urbanismo-infraestructuras/transporte-urbano/poste-autobus?rf=html&srsname=wgs84&start=0&rows=50&distance=500", nil)
	request.Header.Add("Accept", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println("fail!")
		return
	}
	defer resp.Body.Close()
	a, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", a)
}

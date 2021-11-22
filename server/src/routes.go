package src

import (
	"github.com/gorilla/mux"
)

func SetRoutes() *mux.Router {
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/cities", allCities).Methods("GET")
	muxRouter.HandleFunc("/airlines", allAirlines).Methods("GET")
	muxRouter.HandleFunc("/airports", allAirports).Methods("GET")
	muxRouter.HandleFunc("/get_city/{city_name}", getCity).Methods("GET")
	muxRouter.HandleFunc("/get_airline/{airline}", getAirline).Methods("GET")
	muxRouter.HandleFunc("/get_airport/{city_code}", getAirport).Methods("GET")
	return muxRouter
}

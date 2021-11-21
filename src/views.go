package src

import (
	"encoding/json"
	"net/http"
	"strings"
)

func allCities(w http.ResponseWriter, r *http.Request) {
	db := connectDataBase()
	defer db.Close()

	var cities []Cities

	params := r.URL.Query()["search"]
	if len(params) > 0 {
		search := params[0]
		cityName := strings.Title(search)
		db.Where("city LIKE ?", cityName+"%").Or("ru_city LIKE ?", cityName+"%").Find(&cities)
	} else {
		db.Find(&cities)
	}
	json.NewEncoder(w).Encode(cities)
}

func allAirlines(w http.ResponseWriter, r *http.Request) {
	db := connectDataBase()
	defer db.Close()

	var airlines []Airlines
	db.Find(&airlines)

	json.NewEncoder(w).Encode(airlines)
}

func allAirports(w http.ResponseWriter, r *http.Request) {
	db := connectDataBase()
	defer db.Close()

	var airports []Airports
	db.Find(&airports)
	json.NewEncoder(w).Encode(airports)
}

func getCity(w http.ResponseWriter, r *http.Request) {
	db := connectDataBase()
	defer db.Close()

	cityName := getSearchParam(r, "city_name")

	var cities []Cities
	db.Where("city = ?", cityName).Or("ru_city = ?", cityName).Find(&cities)
	json.NewEncoder(w).Encode(cities)
}

func getAirline(w http.ResponseWriter, r *http.Request) {
	db := connectDataBase()
	defer db.Close()

	airlineName := getSearchParam(r, "airline")

	var airlines []Airlines
	db.Where("name LIKE ?", "%"+airlineName+"%").Find(&airlines)
	json.NewEncoder(w).Encode(airlines)
}

func getAirport(w http.ResponseWriter, r *http.Request) {
	db := connectDataBase()
	defer db.Close()

	cityCode := getSearchParam(r, "city_code")

	var airports []Airports
	db.Where("city_code = ?", cityCode).Find(&airports)
	json.NewEncoder(w).Encode(airports)
}

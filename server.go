package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"go-travel/src"
	"net/http"
)

//TODO: 1) Заполнить README.md

func main() {
	fmt.Println("======== Go Travel ========")
	fmt.Println("http://localhost:8081")
	routes := src.SetRoutes()
	http.ListenAndServe(":8081", routes)
}

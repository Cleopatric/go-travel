package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"net/http"
	"server/src"
)

func main() {
	fmt.Println("======== Go Travel ========")
	fmt.Println("http://localhost:8081")
	routes := src.SetRoutes()
	http.ListenAndServe(":8081", routes)
}

package src

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strings"
)

func connectDataBase() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./sqlite.db")
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}
	return db
}

func getSearchParam(r *http.Request, paramName string) string {
	vars := mux.Vars(r)
	param := vars[paramName]
	titleParam := strings.Title(param)
	return titleParam
}

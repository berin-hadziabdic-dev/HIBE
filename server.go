package main

import (
	"net/http"
	"github.com/danc2050/HaveIBeenExploited/AppRoot/searchcontract"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./app/build")))
	http.HandleFunc("/search", searchcontract.SearchContract)
	http.ListenAndServe(":8080", nil)
}


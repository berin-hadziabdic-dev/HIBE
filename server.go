package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./app/build")))
	http.HandleFunc("/search", searchcontract.SearchContract)
	http.ListenAndServe(":8080", nil)
}


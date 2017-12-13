package main

import (
	"log"
	"net/http"
	"restfullapiexample"
)

func main() {

	router := restfullapiexample.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}

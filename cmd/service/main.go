package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	//log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
	//log.Panic().Err(rest.New().Start())
}

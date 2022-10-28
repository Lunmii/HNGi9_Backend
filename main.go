package main

import (
	"fmt"
	"log"
	"net/http"
)

type about struct {
	SlackUsername string `json:"SlackUsername"`
	Backend       bool   `json:"Backend"`
	Age           int    `json:"Age"`
	Bio           string `json:"Bio"`
}
type About []about

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
func main() {
	handleRequests()

}

package main

import (
	"fmt"
	"net/http"
)

type PostRequest struct {
	Data string `json:"data"`
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("GET /solve", getSolve)
	router.HandleFunc("POST /solve", postSolve)

	http.ListenAndServe(":8080", router)

}

func getSolve(w http.ResponseWriter, r *http.Request) {
	fmt.Println("boggle solver endpoint")
}

func postSolve(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post request")
}

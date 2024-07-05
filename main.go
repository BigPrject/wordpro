package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type postRecived struct {
	Board string
}

var wordTrie *Trie = StartTrie()

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
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Could not read request %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var recivedData postRecived
	err = json.Unmarshal(body, &recivedData)
	if err != nil {
		log.Printf("Couldn't unmarshall body")
		w.WriteHeader(http.StatusBadRequest) //400
	}
	log.Printf("board is %v", recivedData.Board)

	boggle := recivedData.Board

	jsonResponse := search(wordTrie, boggle)

	log.Println(jsonResponse)

}

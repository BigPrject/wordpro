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

var wordTrie *Trie = ImportWords()

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/solve", solveHandler)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", router)
}

func solveHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodGet {
		getSolve(w, r)
	} else if r.Method == http.MethodPost {
		postSolve(w, r)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func getSolve(w http.ResponseWriter, r *http.Request) {
	fmt.Println("boggle solver endpoint")
}

func postSolve(w http.ResponseWriter, r *http.Request) {
	fmt.Println("post request")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Could not read request %v", err)
		w.WriteHeader(http.StatusInternalServerError) // HTTP 500
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
	jsonResponse := solve(wordTrie, boggle)
	w.Write(jsonResponse)

}

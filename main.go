package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type IndexResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var port int = 7145

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got / request")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(IndexResponse{Code: 200, Message: "Mirror Operational"})
	})

	log.Printf("Listening on %s\n", strconv.Itoa(port))
	err := http.ListenAndServe(":7145", mux)
	if err != nil {
		log.Fatal(err)
	}
}

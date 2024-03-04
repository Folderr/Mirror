//go:build go1.22

// Mirror relies upon specific updates to the net/http in go 1.22

package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/Folderr/Mirror/user"
)

type IndexResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var port int = 7145

func main() {
	var isInstanceService bool
	var domain string

	flag.BoolVar(&isInstanceService, "service", false, "Whether or not to run Mirror as a service of a Folderr instance")

	flag.StringVar(&domain, "domain", "", "The domain you want to listen on")
	flag.Parse()

	if isInstanceService {
		log.Fatal("Running Mirror as a service of a Folderr instance is not yet supported")
	}
	// if domain == "" {
	//	log.Fatal("Missing the domain flag.\nPlease run \"Mirror-Server -h\" for help")
	// }

	if !isInstanceService && domain != "" {
		err := user.DomainCheck(domain)
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	// handle light config, this is temporary
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//		log.Println("Got / request")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(IndexResponse{Code: 200, Message: "Mirror Operational"})
	})

	log.Printf("Listening on %s\n", strconv.Itoa(port))
	err := http.ListenAndServe(":7145", mux)
	if err != nil {
		log.Fatal(err)
	}
}

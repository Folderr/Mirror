//go:build go1.22

// Mirror relies upon specific updates to the net/http in go 1.22

package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

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

	if isInstanceService {
		log.Fatal("Running Mirror as a service of a Folderr instance is not yet supported")
	}

	flag.StringVar(&domain, "domain", "", "The domain you want to listen on for service mode, or the instance domain in user - default - mode")
	flag.Parse()

	if domain == "" {
		log.Fatal("I need a domain")
	}

	if !strings.Contains(domain, "http://") && !strings.Contains(domain, "https://") {
		if strings.Contains(domain, "://") {
			log.Fatal("I need either a https url or a http url")
		} else {
			domain = "http://" + domain
		}
	}

	// if domain == "" {
	//	log.Fatal("Missing the domain flag.\nPlease run \"Mirror-Server -h\" for help")
	// }

	checkedDomain, err := url.Parse(domain)
	if err != nil {
		log.Fatal("Failed to parse URL", domain, "\nPlease ensure its a valid URL")
	}

	if checkedDomain.Host == "" {
		checkedDomain.Host = checkedDomain.Scheme
		checkedDomain.Scheme = "http"
	}

	// handle light config, this is temporary
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//		log.Println("Got / request")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(IndexResponse{Code: 200, Message: "Mirror Operational"})
	})

	if !isInstanceService {
		user.SetupEndpoints(checkedDomain, mux)
	}

	log.Printf("Listening on %s\n", strconv.Itoa(port))
	err = http.ListenAndServe(":7145", mux)
	if err != nil {
		log.Fatal(err)
	}
}

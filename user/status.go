package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type InfoReturnMessage struct {
	Version      string `json:"version"`
	Node_version string `json:"node_version"`
	Online_since int    `json:"online_since"`
	Message      string `json:"message"`
}

type InfoReturn struct {
	Code    int               `json:"code"`
	Message InfoReturnMessage `json:"message"`
}

func check(domain *url.URL) (isOnline bool, err error) {
	if domain.Scheme == "" {
		domain.Scheme = "http"
	}

	resp, err := http.Get(domain.Scheme + "://" + domain.Host + "/api")
	if err != nil {
		return false, err
	}

	if resp.StatusCode != 200 {
		return false, fmt.Errorf("non 200 status code recieved")
	}

	if resp.Header.Get("Content-Type") == "" || !strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		return false, fmt.Errorf("response type was not json")
	}

	var output InfoReturn
	err = json.NewDecoder(resp.Body).Decode(&output)

	if err != nil {
		return false, err
	}

	if output.Message.Message != "Pong!" && output.Message.Message != "OK" {
		return false, fmt.Errorf("unexpected body return, unsure if Folderr")
	}

	return true, nil
}

var ticker *time.Ticker
var instanceOnline bool

func CheckStatus(domain *url.URL) {
	ticker = time.NewTicker(5 * time.Second)
	instanceOnline = false
	go func() {
		for range ticker.C {
			isOK, err := check(domain)
			urlErrTarget := &url.Error{}
			if errors.As(err, &urlErrTarget) && err.(*url.Error).Timeout() {
				log.Println("Timeout occurred during health check on instance.")
				instanceOnline = false
			} else if errors.As(err, &urlErrTarget) && strings.Contains(err.Error(), "actively refused") {
				log.Println("Health check connection refused. Change!")
				instanceOnline = false
			} else if err != nil {
				fmt.Println("Unknown Error occurred while performing healthcheck, see below")
				log.Println(err)
				instanceOnline = false
			} else {
				instanceOnline = isOK
			}
		}
	}()
}

func StopChecks() {
	if ticker != nil {
		ticker.Stop()
	}
}

package user

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

func check(domain *url.URL) (isOnline bool, rawBody string, err error) {
	if domain.Scheme == "" {
		domain.Scheme = "http"
	}

	resp, err := http.Get(domain.Scheme + "://" + domain.Host + "/api")
	if err != nil {
		return false, "", err
	}

	if resp.StatusCode != 200 {
		return false, "", fmt.Errorf("non 200 status code recieved")
	}

	if resp.Header.Get("Content-Type") == "" || !strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		return false, "", fmt.Errorf("response type was not json")
	}

	bytes, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return false, "", err
	}
	outRaw := string(bytes)

	// let's see if this is valid JSON

	if !json.Valid(bytes) {
		return false, outRaw, fmt.Errorf("response data claimed json. was not json")
	}

	var output InfoReturn
	err = json.Unmarshal(bytes, &output)

	if err != nil && strings.Contains(err.Error(), "cannot unmarshal") && strings.Contains(err.Error(), "type") {
		return false, "", fmt.Errorf("unexpected json output")
	}
	if err != nil {
		return false, outRaw, err
	}

	if output.Message.Message != "Pong!" && output.Message.Message != "OK" {
		return false, outRaw, fmt.Errorf("unexpected body return, unsure if Folderr")
	}

	return true, outRaw, nil
}

var ticker *time.Ticker
var instanceOnline bool

func CheckStatus(domain *url.URL) {
	ticker = time.NewTicker(5 * time.Second)
	instanceOnline = false
	go func() {
		for range ticker.C {
			isOK, rawBody, err := check(domain)
			urlErrTarget := &url.Error{}

			if err != nil {
				if errors.As(err, &urlErrTarget) && err.(*url.Error).Timeout() {
					log.Println("Err: Timeout occurred during health check on instance.")
					instanceOnline = false
				} else if errors.As(err, &urlErrTarget) && strings.Contains(err.Error(), "actively refused") {
					log.Println("Err: Health check connection refused.")
					instanceOnline = false
				} else if err.Error() == "unexpected json output" {
					log.Println("Err: Got back an unexpected body, see below")
					log.Println(rawBody)
				} else if err.Error() == "unexpected body return, unsure if Folderr" {
					log.Println("Not sure this domain is Folderr. Check the body return below (if there is one)")
					if rawBody != "" {
						log.Println(rawBody)
					}
				} else if err.Error() == "response data claimed json. was not json." {
					log.Println("Err (Health Check): response data claimed to be JSON even though it wasn't. See below for data")
					log.Println(rawBody)
				} else {
					log.Println("Unknown Error occurred while performing healthcheck, see below")
					log.Println(err.Error())
				}
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

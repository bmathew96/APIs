package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var requestURL string
var config EnvConfig
var gerr error

func init() {
	requestURL = "https://gateway.marvel.com:443/v1/public"
	config, gerr = getEnvConfig()
	if gerr != nil {
		log.Fatal(gerr)
	}

}

func main() {
	hash, ts := getKeyHash()
	reqURL := fmt.Sprintf("%s/characters?ts=%s&hash=%s&apikey=%s", requestURL, ts, hash, config.APIKeyPublic)

	//send request
	resp, err := http.Get(reqURL)
	if err != nil {
		log.Fatal("Failed to send reuest", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to Read In Body", err)
	}
	// Reset resp.Body so it can be use again
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	fmt.Println("RESPONSE BODY", body)

}

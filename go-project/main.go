package main

import (
	"fmt"
	"log"
	"net/http"
)

var requestURL string
var config EnvConfig
var gerr error

func init() {
	requestURL = "https://gateway.marvel.com:443/v1/public/"
	config, gerr = getEnvConfig()
	if gerr != nil {
		log.Fatal(gerr)
	}

}

func main() {
	hash, ts := getKeyHash()
	reqURL := fmt.Sprintf("%s/characters?ts=%s&hash=%s&apikey=%s", requestURL, ts, hash, config.APIKeyPublic)
	//log.Println("ENV CONFIG: ", config, "ERROR: ", gerr, "HASH:", hash, "TIMESTAMP: ", ts, "RequestString: ", reqURL)

	resp, err := http.Get(reqURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("RESP", resp.Status)

}

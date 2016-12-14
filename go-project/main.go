package main

import (
	"bytes"
	"encoding/json"
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
	//log.Println(config)

}

func main() {
	hash, ts := getKeyHash()
	reqURL := fmt.Sprintf("%s/characters?ts=%s&hash=%s&apikey=%s", requestURL, ts, hash, config.MarvelPublicKey)
	//fmt.Println(reqURL)
	//fmt.Println(config)
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

	var reqResp MarvelResponse
	err = json.Unmarshal(body, &reqResp)
	if err != nil {
		log.Fatal("Failed to umarshall body", err)
	}

	for _, character := range reqResp.Data.Results {
		//	fmt.Println(character.(type))
		switch myMap := character.(type) {
		case []map[string]interface{}:
			for _, v := range myMap {
				tMap := make(map[string]interface{})
				tMap = v
				for a, b := range tMap {
					fmt.Println(a)
					fmt.Println(b)
				}
			}
		}
	}

}

package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// handles the logic to grab all the data and append it.
func getMarvelCharacters() {

	// initial request
	reqURL := getRequestURLString("characters", 100, 0)
	resp, body, err := request(reqURL)

}

// handles the logic to grab all the data and append it.
func getMarvelComics() {

}

//handles all the request to consume the api for us.
func request(reqURL string) (resp *http.Response, body []byte, err error) {

	//send request
	log.Println("\t\t sending request to ", reqURL)
	resp, err = http.Get(reqURL)
	if err != nil {
		err = errors.New("Failed to send reqest: " + err.Error())
		return
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.New("Failed to write body to byte: " + err.Error())
		return
	}
	// Reset resp.Body so it can be use again
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))

	return
}

// formats the request url to our speficication
func getRequestURLString(reqFor string, limit int, offset int) string {
	log.Println("\t\t formatting request url")

	requestURL := "https://gateway.marvel.com:443/v1/public"
	hash, ts := getKeyHash()
	return fmt.Sprintf("%s/%s?ts=%s&hash=%s&apikey=%s&offset=%d&limit=%d", requestURL, reqFor, ts, hash, config.MarvelPublicKey, offset, limit)
}

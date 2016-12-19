package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	env "personal-project/APIs/go-project/environment"
	model "personal-project/APIs/go-project/model"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var config env.EnvConfig
var err error
var db *gorm.DB

func init() {
	log.Println("Getting Environment Config")
	config, err = env.GetEnvConfig()
	if err != nil {
		log.Fatal("Failed: ", err)
	}
	fmt.Println(config)
	log.Println("Connecting to database")
	connectionString := fmt.Sprintf("%s:%s@(%s:%d)/?parseTime=True&loc=%s", config.DBUsername, config.DBPassword, config.DBHostname, config.DBPort, "America%2FChicago")
	db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		log.Fatal("Failed: ", err)
	}
}

func main() {
	defer db.Close()

	status, err := resetCharactersTable()
	log.Println("Status: ", status, "Error: ", err)
	log.Println("At this point we may have sent too much requests ... lets pause for a few minutes")
	time.Sleep(time.Minute * 5)
	log.Println("Resuming ... ")
	status, err = resetComicsTable()
	log.Println("Status: ", status, "Error: ", err)

}

// handles the logic to grab all the data and append it.
func getMarvelCharacters() (characterCollection []model.CharacterTable, err error) {
	log.Println("\t\t Calling all marvel characters")
	response, err := validateAndUnMarshalCharacters(0)
	if err != nil {
		return
	}
	collectionBuilder := model.Characters{}
	characterCollection = collectionBuilder.Populate(response.Data.Results)

	respCount := response.Data.Count
	canSleep := true
	for count := response.Data.Count; count <= response.Data.Total; count += respCount {
		log.Printf("\t\t ON [%d of %d] \n", count, response.Data.Total)

		if count > (response.Data.Total/2) && canSleep {
			log.Println("\t\t We have send too much request ... lets sleep for a miute")
			time.Sleep(time.Minute)
			canSleep = false
			log.Println("\t\t resuming")

		}

		respCount = response.Data.Count
		response, err = validateAndUnMarshalCharacters(count)
		if err != nil {
			return
		}

		characterCollection = collectionBuilder.Populate(response.Data.Results)
	}

	return
}

func resetCharactersTable() (status string, err error) {
	status = "failed to reset characters"
	db.Exec("TRUNCATE TABLE " + model.CharacterTable{}.TableName() + ";")
	collection, err := getMarvelCharacters()
	if err != nil {
		return
	}
	// begin a transaction
	tx := db.Begin()
	log.Println("transaction started")

	charactersCount := len(collection)
	for key, character := range collection {
		log.Printf("adding to queue [%d of %d]", key, charactersCount)
		// add them to queue
		tx.Create(&character)
	}
	// save to table
	tx.Commit()
	log.Println(charactersCount, " rows added")
	status = "characters reset successfully "

	return
}

func validateAndUnMarshalCharacters(offset int) (response model.MarvelResponseCharacter, err error) {
	log.Println("\t\t\t Parsing json to response model")
	reqURL := getRequestURLString("characters", 100, offset)
	resp, body, err := request(reqURL)
	log.Println("\t\t\t Response Status ", resp.StatusCode, resp.Status)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		err = errors.New("Failed Unmarshal: " + err.Error())
		return
	}

	return
}

// handles the logic to grab all the data and append it.
func getMarvelComics() (comicCollection []model.ComicTable, err error) {
	log.Println("\t\t Calling all marvel comics")
	response, err := validateAndUnMarshalComics(0)
	if err != nil {
		return
	}
	collectionBuilder := model.Comics{}
	comicCollection = collectionBuilder.Populate(response.Data.Results)

	respCount := response.Data.Count
	canSleep := true

	for count := response.Data.Count; count <= response.Data.Total; count += respCount {
		log.Printf("\t\t ON [%d of %d] \n", count, response.Data.Total)

		if count > (response.Data.Total/2) && canSleep {
			log.Println("\t\t We have send too much request ... lets sleep for a miute")
			time.Sleep(time.Minute)
			canSleep = false
			log.Println("\t\t resuming")

		}

		respCount = response.Data.Count
		response, err = validateAndUnMarshalComics(count)
		if err != nil {
			return
		}

		comicCollection = collectionBuilder.Populate(response.Data.Results)
	}

	return
}

func resetComicsTable() (status string, err error) {
	status = "failed to reset comic"
	db.Exec("TRUNCATE TABLE " + model.ComicTable{}.TableName() + ";")
	collection, err := getMarvelComics()
	if err != nil {
		return
	}

	// begin a transaction
	tx := db.Begin()
	log.Println("transaction started")

	comicsCount := len(collection)
	for key, comic := range collection {
		log.Printf("adding to queue [%d of %d]", key, comicsCount)
		// add them to queue
		tx.Create(&comic)
	}

	// save to table
	tx.Commit()
	log.Println(comicsCount, " rows added")
	status = "comics reset successfully "

	return
}

func validateAndUnMarshalComics(offset int) (response model.MarvelResponseComic, err error) {
	log.Println("\t\t Parsing json to response model")
	reqURL := getRequestURLString("comics", 100, offset)
	resp, body, err := request(reqURL)
	log.Println("Response Status ", resp.StatusCode, resp.Status)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &response)
	if err != nil {
		err = errors.New("Failed Unmarshal: " + err.Error())
		return
	}

	return
}

//handles all the request to consume the api for us.
func request(reqURL string) (resp *http.Response, body []byte, err error) {

	//send request
	log.Println("\t\t\t sending request ")
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
	log.Println("\t\t\t formatting request url")
	requestURL := "https://gateway.marvel.com:443/v1/public"
	hash, ts := config.GetKeyHash()
	return fmt.Sprintf("%s/%s?ts=%s&hash=%s&apikey=%s&offset=%d&limit=%d", requestURL, reqFor, ts, hash, config.MarvelPublicKey, offset, limit)
}

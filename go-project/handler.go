package main

import (
	"encoding/json"
	"log"
	"net/http"
	"personal-project/APIs/go-project/model"
	"strconv"
	"strings"
)

func getHandlers() {

	http.HandleFunc("/marvel/characters", charactersHandler)
	http.HandleFunc("/marvel/character/", characterHandler)
	http.HandleFunc("/marvel/comics", comicsHandler)
	http.HandleFunc("/marvel/comic/", comicHandler)

}

func charactersHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("\t Method: ", r.Method, "EndPoint: ", r.RequestURI)

	var response []byte
	var resp = model.Response{Code: http.StatusOK, Status: "Netural"}

	searchString := r.FormValue("q")
	limit, err := strconv.Atoi(r.FormValue("l"))
	if err != nil {
		limit = 100
	}

	switch r.Method {
	case "GET":
		log.Println("\t Fetching all characters - Search: ", searchString, "Limit: ", limit)

		resp.Status = "OK"
		resp.Results, err = fetchAllCharacters(searchString, limit)
		if err != nil {
			log.Println("\t Failed: ", err)
		}
		break
	default:
		log.Println("\t Method Not Handled")
		resp.Code = http.StatusServiceUnavailable
		resp.Status = r.Method + " method not handled "
		break

	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(resp.Code)

	response, _ = json.Marshal(resp)
	w.Write(response)

	log.Println("\t ResponseStatus: ", resp.Status, "Response Code: ", resp.Code)

	return
}

func comicsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("\t Method: ", r.Method, "EndPoint: ", r.RequestURI)

	var response []byte
	var resp = model.Response{Code: http.StatusOK, Status: "Netural"}

	searchString := r.FormValue("q")
	limit, err := strconv.Atoi(r.FormValue("l"))
	if err != nil {
		limit = 100
	}

	switch r.Method {
	case "GET":
		log.Println("\t Fetching all comics - Search: ", searchString, "Limit: ", limit)

		resp.Status = "OK"
		resp.Results = fetchAllComics(searchString, limit)
		break
	default:
		resp.Code = http.StatusServiceUnavailable
		resp.Status = r.Method + " method not handled "
		break

	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(resp.Code)

	response, _ = json.Marshal(resp)
	w.Write(response)

	log.Println("\t ResponseStatus: ", resp.Status, "Response Code: ", resp.Code)

	return
}

func characterHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("\t Method: ", r.Method, "EndPoint: ", r.RequestURI)

	var response []byte
	var resp = model.Response{Code: http.StatusOK, Status: "Netural"}

	// 0 - empty 1 - marvel 2 - character 3 - ID
	paths := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(strings.Split(paths[3], "?")[0])
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Status = err.Error()

	}

	switch r.Method {
	case "GET":
		log.Println("\t Fetching Character - ID: ", id)

		if resp.Code == http.StatusOK {
			resp.Status = "Not Vaild ID "

			if id >= 1 {
				resp.Status = "OK"
				resp.Results = fetchCharacter(id)
			}

		}
		break
	case "POST":
		log.Println("\t Updating/Adding Character - ID: ", id)

		var request model.CharacterTable
		err = json.NewDecoder(r.Body).Decode(&request)

		if err == nil {
			request.CharacterID = id

			resp.Code = http.StatusCreated
			resp.Status = "Created"
			if id > 0 {
				resp.Code = http.StatusOK
				resp.Status = "Updated"
			}

			resp.Results = addUpdateCharacter(request)
		}
		break

	case "PUT":
		log.Println("\t Updating/Adding Character - ID: ", id)

		var request model.CharacterTable
		err = json.NewDecoder(r.Body).Decode(&request)

		if err == nil {
			request.CharacterID = id

			resp.Code = http.StatusCreated
			resp.Status = "Created"
			if id > 0 {
				resp.Code = http.StatusOK
				resp.Status = "Updated"
			}

			resp.Results = addUpdateCharacter(request)
		}

		break
	case "DELETE":
		log.Println("\t Removing Character - ID: ", id)
		if resp.Code == http.StatusOK {
			resp.Status = "Not Vaild ID "

			if id >= 1 {
				deleteCharacter(id)
				resp.Status = "DELETED"
			}

		}
		break
	default:
		resp.Code = http.StatusServiceUnavailable
		resp.Status = r.Method + " method not handled "
		break

	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(resp.Code)

	response, _ = json.Marshal(resp)
	w.Write(response)

	log.Println("\t ResponseStatus: ", resp.Status, "Response Code: ", resp.Code)

	return
}

func comicHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("\t Method: ", r.Method, "EndPoint: ", r.RequestURI)

	var response []byte
	var resp = model.Response{Code: http.StatusOK, Status: "Netural"}

	// 0 - empty 1 - marvel 2 - comic 3 - ID
	paths := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(strings.Split(paths[3], "?")[0])
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Status = err.Error()

	}

	switch r.Method {
	case "GET":
		log.Println("\t Fetching Comic - ID: ", id)
		if resp.Code == http.StatusOK {
			resp.Status = "Not Vaild ID "

			if id >= 1 {
				resp.Status = "OK"
				resp.Results = fetchComic(id)
			}

		}
		break
	case "POST":
		log.Println("\t Adding/Updating Comic - ID: ", id)

		var request model.ComicTable
		err = json.NewDecoder(r.Body).Decode(&request)

		if err == nil {
			request.ComicID = id

			resp.Code = http.StatusCreated
			resp.Status = "Created"
			if id > 0 {
				resp.Code = http.StatusOK
				resp.Status = "Updated"
			}

			resp.Results = addUpdateComic(request)
		}
	case "PUT":
		log.Println("\t Adding/Updating Comic - ID: ", id)

		var request model.ComicTable
		err = json.NewDecoder(r.Body).Decode(&request)

		if err == nil {
			request.ComicID = id

			resp.Code = http.StatusCreated
			resp.Status = "Created"
			if id > 0 {
				resp.Code = http.StatusOK
				resp.Status = "Updated"
			}

			resp.Results = addUpdateComic(request)
		}

		break
	case "DELETE":
		log.Println("\t Removing Comic - ID: ", id)
		if resp.Code == http.StatusOK {
			resp.Status = "Not Vaild ID "

			if id >= 1 {
				deleteComic(id)
				resp.Status = "DELETED"
			}

		}
		break
	default:
		resp.Code = http.StatusServiceUnavailable
		resp.Status = r.Method + " method not handled "
		break

	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(resp.Code)

	response, _ = json.Marshal(resp)
	w.Write(response)

	log.Println("\t ResponseStatus: ", resp.Status, "Response Code: ", resp.Code)

	return
}

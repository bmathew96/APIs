package main

import (
	"encoding/json"
	"fmt"
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
	var response []byte
	var resp = model.Response{Code: http.StatusOK, Status: "Netural"}

	searchString := r.FormValue("q")

	switch r.Method {
	case "GET":
		resp.Status = "OK"
		resp.Results = fetchAllCharacters(searchString)
		break
	default:
		resp.Code = http.StatusServiceUnavailable
		resp.Status = r.Method + " method not handled "
		break

	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(resp.Code)

	response, _ = json.Marshal(resp)
	w.Write(response)

	return
}

func comicsHandler(w http.ResponseWriter, r *http.Request) {
	var response []byte
	var resp = model.Response{Code: http.StatusOK, Status: "Netural"}

	searchString := r.FormValue("q")

	switch r.Method {
	case "GET":
		resp.Status = "OK"
		resp.Results = fetchAllComics(searchString)
		break
	default:
		resp.Code = http.StatusServiceUnavailable
		resp.Status = r.Method + " method not handled "
		break

	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(resp.Code)

	response, _ = json.Marshal(resp)
	w.Write(response)

	return
}

func characterHandler(w http.ResponseWriter, r *http.Request) {
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
		if resp.Code == http.StatusOK {
			resp.Status = "Not Vaild ID "
			fmt.Println("ID: ", id)
			if id >= 1 {
				resp.Status = "OK"
				resp.Results = fetchCharacter(id)
			}

		}
		break
	case "POST":
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
		if resp.Code == http.StatusOK {
			resp.Status = "Not Vaild ID "
			fmt.Println("ID: ", id)
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
	w.WriteHeader(resp.Code)

	response, _ = json.Marshal(resp)
	w.Write(response)

	return
}

func comicHandler(w http.ResponseWriter, r *http.Request) {
	var response []byte
	var resp = model.Response{Code: http.StatusOK, Status: "Netural"}

	// 0 - empty 1 - marvel 2 - comic 3 - ID
	paths := strings.Split(r.RequestURI, "/")
	id, err := strconv.Atoi(strings.Split(paths[3], "?")[0])
	if err != nil {
		resp.Code = http.StatusBadRequest
		resp.Status = err.Error()

	}
	fmt.Println("Method: ", r.Method)
	switch r.Method {
	case "GET":
		if resp.Code == http.StatusOK {
			resp.Status = "Not Vaild ID "
			fmt.Println("ID: ", id)
			if id >= 1 {
				resp.Status = "OK"
				resp.Results = fetchComic(id)
			}

		}
		break
	case "POST":
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
		if resp.Code == http.StatusOK {
			resp.Status = "Not Vaild ID "
			fmt.Println("ID: ", id)
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
	w.WriteHeader(resp.Code)

	response, _ = json.Marshal(resp)
	w.Write(response)

	return
}

package main

// func charactersHandler(w http.ResponseWriter, r *http.Request) {

// 	paths := strings.Split(r.RequestURI, "/")
// 	var response []byte
// 	var resp model.Response
// 	resp.Code = http.StatusOK
// 	resp.Status = "Netural"
// 	fmt.Println("PATHS: ", paths)
// 	fmt.Println("Query", r.FormValue("q"))
// 	// status, err := helper.restCharactersTable()
// 	// if err != nil {
// 	// 	resp.Code = http.StatusInternalServerError
// 	// 	resp.Status = status + " - " + err.Error()
// 	// }
// 	//resp.Status = status

// 	w.Header().Set("Content-Type", "application/json; charset=utf-8")
// 	w.WriteHeader(resp.Code)

// 	response, _ = json.Marshal(resp)
// 	w.Write(response)

// 	return
// }

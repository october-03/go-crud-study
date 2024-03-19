package handlers

import (
	"crud-study/crud-study/dto"
	"crud-study/crud-study/services"
	"crud-study/crud-study/utils"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	urlParts := strings.Split(path, "/")

	body, err := io.ReadAll(r.Body)

	if err != nil {
		panic(err)
	}

	endPoint := urlParts[2]

	//urlParams := urlParts[3]
	//urlQuery := r.URL.Query()

	fmt.Println(endPoint)

	defer func() {
		if r := recover(); r != nil {
			data, _ := json.Marshal(r)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(data)
			return
		}
	}()

	switch endPoint {
	case "signup":
		if r.Method != "POST" {
			utils.MethodNotAllowed(w)
			return
		}

		var data dto.CreateUserDto
		if err := json.Unmarshal(body, &data); err != nil {
			panic(err)
		}

		res := services.CreateUser(data)
		w.Write(res)
	default:
		fmt.Println(path + " No endpoint")
		http.NotFound(w, r)
	}
}

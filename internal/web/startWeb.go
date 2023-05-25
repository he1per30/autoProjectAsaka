package web

import (
	"autoProjectAsaka/internal"
	"autoProjectAsaka/internal/domain"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

type SomeInterface interface {
}

func StartWeb() {
	log.Println("Starting the HTTP server on port 8080")
	router := mux.NewRouter()

	router.HandleFunc("/api/createApp", postCreateAdapter).Methods("POST")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln(err)
	}

}

func postCreateAdapter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var app domain.AppAsaka
	err = json.Unmarshal(reqBody, &app)
	if err != nil {
		fmt.Println(err)
		return
	}
	internal.ManageFile(app)
}

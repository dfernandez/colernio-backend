package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Course struct {
	Id   int
	Name string
}

var Router = func() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index).Methods("GET")

	return router
}()

func New() {
	log.Fatal(http.ListenAndServe(":8080", Router))
}

func Index(w http.ResponseWriter, r *http.Request) {

	courses := make(map[string]Course)
	courses["course:1"] = Course{1, "Hello world 101"}
	courses["course:2"] = Course{2, "Hello world intermediate"}
	courses["course:3"] = Course{3, "Hello world intermediate"}
	courses["course:4"] = Course{4, "Hello world intermediate"}
	courses["course:5"] = Course{5, "Hello world intermediate"}
	courses["course:6"] = Course{6, "Hello world intermediate"}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(courses)
}

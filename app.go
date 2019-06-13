package main

import (
	. "SmashLocalFinder/config"
	. "SmashLocalFinder/dao"
	. "SmashLocalFinder/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var dao = LocalsDAO{}
var config = Config{}

func AllLocals(w http.ResponseWriter, r *http.Request) {
	locals, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, locals)
}

func FindLocal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func CreateLocal(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var local Local
	if err := json.NewDecoder(r.Body).Decode(&local); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	local.ID = bson.NewObjectId()
	if err := dao.Insert(local); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, local)
}

func UpdateLocal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func DeleteLocal(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init() {
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/locals", AllLocals).Methods("GET")
	r.HandleFunc("/locals", CreateLocal).Methods("POST")
	r.HandleFunc("/locals", UpdateLocal).Methods("PUT")
	r.HandleFunc("/locals", DeleteLocal).Methods("DELETE")
	r.HandleFunc("/locals/{id}", FindLocal).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

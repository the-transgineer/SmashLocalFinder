package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	. "github.com/the-transgineer/SmashLocalFinder/config"
	. "github.com/the-transgineer/SmashLocalFinder/dao"
	. "github.com/the-transgineer/SmashLocalFinder/models"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var dao = LocalsDAO{}
var config = Config{}

//enableCors enables CORS
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

//AllLocals retuns a list of all Smash Ultimate Locals
func AllLocals(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	locals, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, locals)
}

//FindLocal Returns a single local based on an ID
func FindLocal(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	local, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}
	respondWithJSON(w, http.StatusOK, local)
}

//FindLocalsByRegion returns all locals in a region
func FindLocalsByRegion(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	locals, err := dao.FindByRegion(params["region"])
	if err != nil {
		return
	}
	respondWithJSON(w, http.StatusOK, locals)

}

//CreateLocal creates a new Local in the DB
func CreateLocal(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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
	respondWithJSON(w, http.StatusCreated, local)
}

//UpdateLocal updates existing Local
func UpdateLocal(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	defer r.Body.Close()
	var local Local
	if err := json.NewDecoder(r.Body).Decode(&local); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"results": "success"})
}

//DeleteLocal deletes local by ID
func DeleteLocal(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	defer r.Body.Close()
	fmt.Println(r.Body)
	var local Local
	if err := json.NewDecoder(r.Body).Decode(&local); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid payload")
		return
	}
	if err := dao.Delete(local); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

//respondWithError is a helper function to send an error
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

//respondWithJSON is a helper functino to send JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
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
	r.HandleFunc("/locals/regions/{region}", FindLocalsByRegion).Methods("GET")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

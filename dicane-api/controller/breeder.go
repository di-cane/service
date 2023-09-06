package controller

import (
	"dicane-api/data"
	"dicane-api/helpers"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var breeder data.Breeder

func GetBreederByEmail(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "email")
	breeder, err := breeder.GetByEmail(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, breeder)
}

func InsertBreeder(w http.ResponseWriter, r *http.Request) {
	var breeder data.Breeder

	err := json.NewDecoder(r.Body).Decode(&breeder)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	breederInserted, err := breeder.Insert(breeder)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, breederInserted)
}

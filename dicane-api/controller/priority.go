package controller

import (
	"dicane-api/data"
	"dicane-api/helpers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var priority data.Priority

func GetPriorityList(w http.ResponseWriter, r *http.Request) {
	saleId := chi.URLParam(r, "sale_id")
	priority, err := priority.GetPriorityList(saleId)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, priority)
}

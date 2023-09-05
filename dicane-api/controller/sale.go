package controller

import (
	"dicane-api/data"
	"dicane-api/helpers"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var sale data.Sale

func GetAllSales(w http.ResponseWriter, r *http.Request) {
	all, err := sale.GetAll()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, all)
}

// GET//sale/{id}
func GetSaleById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	sale, err := sale.GetOne(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, sale)
}

func InsertSale(w http.ResponseWriter, r *http.Request) {
	var sale data.Sale

	err := json.NewDecoder(r.Body).Decode(&sale)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	saleInserted, err := sale.Insert(sale)

	// CHECK
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, saleInserted)
}

// PUT/sale/{id}
func UpdateSale(w http.ResponseWriter, r *http.Request) {
	var saleData data.Sale
	id := chi.URLParam(r, "id")
	err := json.NewDecoder(r.Body).Decode(&saleData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	saleUpdated, err := sale.Update(id, saleData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, saleUpdated)
}

// DELETE/sale/{id}
func DeleteSale(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := sale.DeleteByID(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"message": "successfull deletion"})
}

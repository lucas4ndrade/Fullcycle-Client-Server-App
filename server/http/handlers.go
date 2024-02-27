package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-fullcicle/server/database"
	"github.com/go-fullcicle/server/services"
)

func getBid(w http.ResponseWriter, r *http.Request) {
	bid, err := services.GetDollarBid()
	if err != nil {
		logError(err)
		respondJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	err = database.SaveBid(bid)
	if err != nil {
		logError(err)
		respondJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(w, http.StatusOK, bid)
}

// respondJSON makes the response with a payload in json format
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_, _ = w.Write([]byte(response))
}

// logError logs a error to stdout
func logError(err error) {
	fmt.Printf("Failed to process request, %v", err)
}

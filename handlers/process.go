package handlers

import (
	"log"
	"net/http"
	"strconv"

	"go-league-challenge/models"
	"go-league-challenge/utils"
)

const (
	actionEcho     = "echo"
	actionInvert   = "invert"
	actionFlatten  = "flatten"
	actionSum      = "sum"
	actionMultiply = "multiply"
)

// HandleProcess processes the CSV file based on the action specified in the query parameters
func HandleProcess(w http.ResponseWriter, r *http.Request) {
	action := r.URL.Query().Get("action")
	records, err := utils.ParseCSVFile(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error parsing CSV: %v", err)
		return
	}

	var response string
	switch action {
	case actionEcho:
		response = models.FormatRecords(records)
	case actionInvert:
		response = models.FormatRecords(models.InvertMatrix(records))
	case actionFlatten:
		response = models.FlattenMatrix(records)
	case actionSum:
		response = strconv.Itoa(models.SumMatrix(records))
	case actionMultiply:
		response = strconv.Itoa(models.MultiplyMatrix(records))
	default:
		http.Error(w, "Invalid action", http.StatusBadRequest)
		log.Printf("Invalid action: %s", action)
		return
	}

	_, err = w.Write([]byte(response))
	if err != nil {
		log.Printf("Error writing response: %v", err)
	}
}

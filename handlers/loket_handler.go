package handlers

import (
	"encoding/json"
	"jm-CICO/services"
	"jm-CICO/utils"
	"net/http"
)

func LoketCabangHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}
	url := "https://docs.google.com/spreadsheets/d/14qJDOG1KR8_BxOMgDbgxosnXtl-45enkNT1e3yXgmTU/gviz/tq?tqx=out:csv&sheet=1%20LOK%20CAB"
	data, err := services.FetchLoketData(url)
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"data": data})
}
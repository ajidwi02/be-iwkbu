package handlers

import (
	"encoding/json"
	"jm-CICO/services"
	"jm-CICO/utils"
	"net/http"
)

func SamsatPurwantoroHandler(w http.ResponseWriter, r *http.Request){
	utils.EnableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}
	url := "https://docs.google.com/spreadsheets/d/14qJDOG1KR8_BxOMgDbgxosnXtl-45enkNT1e3yXgmTU/gviz/tq?tqx=out:csv&sheet=54%20Purwantoro"
	data, err := services.FetchPurwantoroData(url)
	if err != nil {
		http.Error(w, "Gagal mengambil data dari sheet Purwantoro: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"data": data})
}
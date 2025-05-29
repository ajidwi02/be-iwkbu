package handlers

import (
	"encoding/json"
	"jm-CICO/services"
	"jm-CICO/utils"
	"net/http"
)

func SamsatLokPrwSKHHandler(w http.ResponseWriter, r *http.Request){
	utils.EnableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}
	url := "https://docs.google.com/spreadsheets/d/14qJDOG1KR8_BxOMgDbgxosnXtl-45enkNT1e3yXgmTU/gviz/tq?tqx=out:csv&sheet=50%20Lok%20Prw%20SKH"
	data, err := services.FetchLokPrwSKHData(url)
	if err != nil {
		http.Error(w, "Gagal mengambil data dari sheet LokPrwSKH: "+err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{"data": data})
}
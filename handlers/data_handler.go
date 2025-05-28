package handlers

import (
	"encoding/json"
	"jm-CICO/config"
	"jm-CICO/models"
	"jm-CICO/services"
	"net/http"
	"strings"
)

func LocationHandlerFactory(cfg config.LocationConfig) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		// ekstrak parameter path
		pathParts := strings.Split(r.URL.Path, "/")
		isSubtotal := len(pathParts) > 4 && pathParts[4] == "subtotal"

		data, err := services.FetchData(cfg)
		if err != nil {
			http.Error(w, "failed to fetch data: "+err.Error(),http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")

		if isSubtotal {
			subtotal := services.CalculateSubtotal(data)
			response := struct {
				Data []models.Entry `json:"data"`
				Subtotal models.Subtotal `json:"subtotal"`
			}{
				Data: data,
				Subtotal: subtotal,
			}
			json.NewEncoder(w).Encode(response)
		} else {
			json.NewEncoder(w).Encode(map[string][]models.Entry{"data": data})
		}
	}
}
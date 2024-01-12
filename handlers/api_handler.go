// handlers/api_handler.go

package handlers

import (
	"fmt"
	"net/http"
	"text/template"
	"time"

	stations "main.go/Stations_1.15.2/stations"
)

// APIHandler обрабатывает запросы на странице API.
func APIHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		renderAPITemplate(w)
	case http.MethodPost:
		handleAPISubmit(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func renderAPITemplate(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("templates/api.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func handleAPISubmit(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	stationID := r.FormValue("stationID")

	stationsClient, err := stations.NewStationsClient("https://apis.deutschebahn.com/db-api-marketplace/apis/ris-stations/v1", "4c3a007856e1d26ef8eb46d25bb5369c", "04581570425fb991801ee47ac31de580", time.Second*10)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	station, err := stationsClient.GetStation(stationID)
	if err != nil {
		http.Error(w, fmt.Sprintf("GET request failed: %v", err), http.StatusInternalServerError)
		return
	}

	renderAPITemplateWithStation(w, station)
}

func renderAPITemplateWithStation(w http.ResponseWriter, station stations.Station) {
	tmpl, err := template.ParseFiles("templates/api.html")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, station); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

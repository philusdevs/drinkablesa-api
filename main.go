package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Municipality struct {
	Name               string `json:"name"`
	Municipal          string `json:"municipal"`
	Disinfectant       string `json:"disinfectant"`
	NonHealthAesthetic string `json:"non_health_aesthetic"`
}

type MunicipalityData struct {
	Municipalities []Municipality `json:"municipalities"`
}

func main() {
	file, err := os.Open("data.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var data MunicipalityData
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	http.HandleFunc("/municipalities/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		name := r.URL.Path[len("/municipalities/"):]
		decodedName, err := url.QueryUnescape(name)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Invalid municipality name"})
			return
		}

		for _, municipality := range data.Municipalities {
			if municipality.Name == decodedName {
				json.NewEncoder(w).Encode(municipality)
				return
			}
		}

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Municipality not found"})
	})

	log.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

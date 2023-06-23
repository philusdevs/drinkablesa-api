package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback to port 8080 if $PORT is not set
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

	log.Println("Server started on http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

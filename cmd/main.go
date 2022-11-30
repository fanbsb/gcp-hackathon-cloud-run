package main

import (
	"encoding/json"
	"gcp-hackathon-cloud-run/api"
	"gcp-hackathon-cloud-run/types"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	port := "8080"
	if v := os.Getenv("PORT"); v != "" {
		port = v
	}
	http.HandleFunc("/", handler)
	http.HandleFunc("/ping", ping)

	log.Printf("starting server on port :%s", port)
	err := http.ListenAndServe(":"+port, nil)
	log.Fatalf("http listen error: %v", err)
}

func handler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		_, _ = io.WriteString(w, "Let the battle begin!")
		return
	}

	var reqBody types.ArenaUpdate
	defer req.Body.Close()
	d := json.NewDecoder(req.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&reqBody); err != nil {
		log.Printf("WARN: failed to decode ArenaUpdate in response body: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	resp := api.Play(reqBody)
	_, _ = io.WriteString(w, resp)
}

func ping(w http.ResponseWriter, req *http.Request) {
	hostname, _ := os.Hostname()
	_, _ = io.WriteString(w, hostname)
}

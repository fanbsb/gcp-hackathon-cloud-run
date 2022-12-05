package main

import (
	"encoding/json"
	"gcp-hackathon-cloud-run/api"
	"gcp-hackathon-cloud-run/types"
	"log"
	"os"
	"testing"
)

func TestJsonUnmarshal(t *testing.T) {
	file, _ := os.Open("/Users/chainge1/GolandProjects/gcp-hackathon-cloud-run/sample/request.json")

	var reqBody types.ArenaUpdate
	d := json.NewDecoder(file)
	d.DisallowUnknownFields()
	if err := d.Decode(&reqBody); err != nil {
		log.Fatal(err)
	}

	log.Println(reqBody.Arena.Dimensions)

	api.Play(reqBody)
}

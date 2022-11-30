package api

import (
	"encoding/json"
	"gcp-hackathon-cloud-run/types"
	"log"
)

const (
	FORWARD = "F"
	LEFT    = "L"
	RIGHT   = "R"
	THROW   = "T"
)

var playground = [][]string{}

func Play(input types.ArenaUpdate) string {
	jstr, _ := json.Marshal(input)
	log.Println(string(jstr))

	commands := []string{"F", "R", "L", "T"}
	//rand := rand2.Intn(4)

	// TODO add your implementation here to replace the random response
	return commands[0]
}

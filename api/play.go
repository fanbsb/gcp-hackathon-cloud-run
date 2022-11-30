package api

import (
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
	log.Printf("IN: %#v", input)

	commands := []string{"F", "R", "L", "T"}
	//rand := rand2.Intn(4)

	// TODO add your implementation here to replace the random response
	return commands[0]
}

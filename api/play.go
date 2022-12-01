package api

import (
	"gcp-hackathon-cloud-run/types"
)

const (
	// Direction
	NORTH = "N"
	SOUTH = "S"
	WEST  = "W"
	EAST  = "E"

	// Command
	FORWARD = "F"
	LEFT    = "L"
	RIGHT   = "R"
	THROW   = "T"
)

var playground [][]bool

func Play(input types.ArenaUpdate) string {
	regroup(input.Arena.Dimensions, input.Arena.State)

	return action(input.Arena.State[input.Links.Self.Href])
}

func regroup(ground []int, states map[string]types.PlayerState) {
	// init the playground
	playground = make([][]bool, ground[0])
	for i := 0; i < ground[0]; i++ {
		playground[i] = make([]bool, ground[1])
	}

	// fill the playground
	for _, v := range states {
		playground[v.X][v.Y] = true
	}
}

func action(self types.PlayerState) string {
	// Throw
	if throw(self) {
		return THROW
	}
	// Runaway
	if self.WasHit {
		return wasHit(self)
	}

	return findPlayerNearBy(self)
}

// Find the player and throw
func throw(self types.PlayerState) bool {
	switch self.Direction {
	case NORTH:
		return hasPlayer(self.X, self.Y-1)
	case SOUTH:
		return hasPlayer(self.X, self.Y+1)
	case WEST:
		return hasPlayer(self.X-1, self.Y)
	case EAST:
		return hasPlayer(self.X+1, self.Y)
	}
	return false
}

// Find the player and runaway
func wasHit(self types.PlayerState) string {
	switch self.Direction {
	case NORTH:
		return hasBlockToMove(self.X, self.Y-1)
	case SOUTH:
		return hasBlockToMove(self.X, self.Y+1)
	case WEST:
		return hasBlockToMove(self.X-1, self.Y)
	case EAST:
		return hasBlockToMove(self.X+1, self.Y)
	}
	return ""
}

// Find Player NearBy
func findPlayerNearBy(self types.PlayerState) string {
	//
	return FORWARD
}

func hasPlayer(x, y int) bool {
	if x < 0 || y < 0 || x >= len(playground) || y >= len(playground[0]) {
		return false
	}
	return playground[x][y]
}

func hasBlockToMove(x, y int) string {
	if x < 0 || y < 0 || x >= len(playground) || y >= len(playground[0]) {
		return LEFT
	}
	if !playground[x][y] {
		return FORWARD
	} else {
		return RIGHT
	}
}

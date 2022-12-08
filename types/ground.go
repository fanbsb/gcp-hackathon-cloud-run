package types

import (
	"log"
	"math"
)

type Ground struct {
	field [][]*int

	highestScoreX, highestScoreY int
}

// InitGround init the playground
func InitGround(ground []int, states map[string]PlayerState) *Ground {
	field := make([][]*int, ground[0])
	for i := 0; i < ground[0]; i++ {
		field[i] = make([]*int, ground[1])
	}

	highestScore := 0
	var highestX, highestY int
	// fill the playground
	for _, v := range states {
		player := v
		field[player.X][player.Y] = &player.Score

		if player.Score >= highestScore {
			highestX, highestY = player.X, player.Y
		}
	}
	return &Ground{
		field:         field,
		highestScoreX: highestX,
		highestScoreY: highestY,
	}
}

func (g *Ground) northPlayer(x, y int) (bool, int) {
	if has, score := g.hasPlayer(x, y-1); has {
		return true, score
	} else if has, score := g.hasPlayer(x, y-2); has {
		return true, score
	} else {
		return g.hasPlayer(x, y-3)
	}
}

func (g *Ground) southPlayer(x, y int) (bool, int) {
	if has, score := g.hasPlayer(x, y+1); has {
		return true, score
	} else if has, score := g.hasPlayer(x, y+2); has {
		return true, score
	} else {
		return g.hasPlayer(x, y+3)
	}
}

func (g *Ground) westPlayer(x, y int) (bool, int) {
	if has, score := g.hasPlayer(x-1, y); has {
		return true, score
	} else if has, score := g.hasPlayer(x-2, y); has {
		return true, score
	} else {
		return g.hasPlayer(x-3, y)
	}
}

func (g *Ground) eastPlayer(x, y int) (bool, int) {
	if has, score := g.hasPlayer(x+1, y); has {
		return true, score
	} else if has, score := g.hasPlayer(x+2, y); has {
		return true, score
	} else {
		return g.hasPlayer(x+3, y)
	}
}

func (g *Ground) hasPlayer(x, y int) (bool, int) {
	if x < 0 || y < 0 || x >= len(g.field) || y >= len(g.field[0]) {
		return false, math.MinInt
	}

	if score := g.field[x][y]; score != nil {
		return true, *score
	}
	return false, math.MinInt
}

func (g *Ground) isNotEdge(x, y int) bool {
	if x < 0 || y < 0 || x >= len(g.field) || y >= len(g.field[0]) {
		return false
	}
	return true
}

func (g *Ground) hasBlockToMove(x, y int) bool {
	return g.isNotEdge(x, y) && g.field[x][y] == nil
}

func (g *Ground) HighestPlayer() (int, int) {
	return g.highestScoreX, g.highestScoreY
}

// Print For debug purpose, printout the ground matrix
func (g Ground) Print() {
	for i := 0; i < len(g.field); i++ {
		log.Printf("%v\n", g.field[i])
	}
	log.Println()
}

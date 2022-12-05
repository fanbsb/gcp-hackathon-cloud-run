package types

import (
	"log"
	"math"
)

type Ground struct {
	field [][]*int
}

// InitGround init the playground
func InitGround(ground []int, states map[string]PlayerState) *Ground {
	field := make([][]*int, ground[0])
	for i := 0; i < ground[0]; i++ {
		field[i] = make([]*int, ground[1])
	}

	// fill the playground
	for _, v := range states {
		field[v.X][v.Y] = &v.Score
	}
	return &Ground{
		field: field,
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

func (g *Ground) hasBlockToMove(x, y int) bool {
	if x < 0 || y < 0 || x >= len(g.field) || y >= len(g.field[0]) {
		return false
	}
	return g.field[x][y] == nil
}

func (g Ground) Print() {
	for i := 0; i < len(g.field); i++ {
		log.Printf("%d\n", g.field[i])
	}
	log.Println()
}

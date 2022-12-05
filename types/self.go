package types

import "math"

type Myself struct {
	x, y      int
	direction string
	wasHit    bool
}

// InitMyself init the object
func InitMyself(x, y int, direction string, wasHit bool) *Myself {
	return &Myself{
		x:         x,
		y:         y,
		direction: direction,
		wasHit:    wasHit,
	}
}

func (m *Myself) WasHit() bool {
	return m.wasHit
}

func (m *Myself) FrontHasPlayer(ground *Ground) (bool, int) {
	switch m.direction {
	case NORTH:
		return ground.northPlayer(m.x, m.y)
	case SOUTH:
		return ground.southPlayer(m.x, m.y)
	case WEST:
		return ground.westPlayer(m.x, m.y)
	case EAST:
		return ground.eastPlayer(m.x, m.y)
	}
	return false, math.MinInt
}

func (m *Myself) BackHasPlayer(ground *Ground) (bool, int) {
	switch m.direction {
	case NORTH:
		return ground.southPlayer(m.x, m.y)
	case SOUTH:
		return ground.northPlayer(m.x, m.y)
	case WEST:
		return ground.eastPlayer(m.x, m.y)
	case EAST:
		return ground.westPlayer(m.x, m.y)
	}
	return false, math.MinInt
}

func (m *Myself) LeftHasPlayer(ground *Ground) (bool, int) {
	switch m.direction {
	case NORTH:
		return ground.westPlayer(m.x, m.y)
	case SOUTH:
		return ground.eastPlayer(m.x, m.y)
	case WEST:
		return ground.southPlayer(m.x, m.y)
	case EAST:
		return ground.northPlayer(m.x, m.y)
	}
	return false, math.MinInt
}

func (m *Myself) RightHasPlayer(ground *Ground) (bool, int) {
	switch m.direction {
	case NORTH:
		return ground.eastPlayer(m.x, m.y)
	case SOUTH:
		return ground.westPlayer(m.x, m.y)
	case WEST:
		return ground.northPlayer(m.x, m.y)
	case EAST:
		return ground.southPlayer(m.x, m.y)
	}
	return false, math.MinInt
}

func (m *Myself) MoveFront(ground *Ground) string {
	if m.canMoveFront(ground) {
		return FORWARD
	} else if m.canMoveLeft(ground) {
		return LEFT
	} else if m.canMoveRight(ground) {
		return RIGHT
	}
	return RIGHT
}

func (m *Myself) MoveLeft(ground *Ground) string {
	if m.canMoveLeft(ground) {
		return LEFT
	} else if m.canMoveRight(ground) {
		return RIGHT
	} else if m.canMoveFront(ground) {
		return FORWARD
	}
	return RIGHT
}

func (m *Myself) MoveRight(ground *Ground) string {
	if m.canMoveRight(ground) {
		return RIGHT
	} else if m.canMoveLeft(ground) {
		return LEFT
	} else if m.canMoveFront(ground) {
		return FORWARD
	}
	return RIGHT
}

func (m *Myself) canMoveFront(ground *Ground) bool {
	switch m.direction {
	case NORTH:
		return ground.hasBlockToMove(m.x, m.y-1)
	case SOUTH:
		return ground.hasBlockToMove(m.x, m.y+1)
	case WEST:
		return ground.hasBlockToMove(m.x-1, m.y)
	case EAST:
		return ground.hasBlockToMove(m.x+1, m.y)
	}
	return false
}

func (m *Myself) canMoveLeft(ground *Ground) bool {
	switch m.direction {
	case NORTH:
		return ground.hasBlockToMove(m.x-1, m.y)
	case SOUTH:
		return ground.hasBlockToMove(m.x+1, m.y)
	case WEST:
		return ground.hasBlockToMove(m.x, m.y+1)
	case EAST:
		return ground.hasBlockToMove(m.x, m.y-1)
	}
	return false
}

func (m *Myself) canMoveRight(ground *Ground) bool {
	switch m.direction {
	case NORTH:
		return ground.hasBlockToMove(m.x+1, m.y)
	case SOUTH:
		return ground.hasBlockToMove(m.x-1, m.y)
	case WEST:
		return ground.hasBlockToMove(m.x, m.y-1)
	case EAST:
		return ground.hasBlockToMove(m.x, m.y+1)
	}
	return false
}

func (m *Myself) HigherScorePlayer(ground *Ground) {

}

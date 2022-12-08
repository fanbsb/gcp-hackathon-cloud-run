package types

import (
	"math"
	"sort"
)

type Myself struct {
	x, y      int
	direction string
	wasHit    bool

	enemyL, enemyR, enemyF, enemyB bool
	highestScoreAroundMe           int
}

type position struct {
	position int
	score    int
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

func (m *Myself) InitMyselfInGround(ground *Ground) {
	// Position
	var positionSlice []position
	fHas, fScore := m.frontHasPlayer(ground)
	lHas, lScore := m.leftHasPlayer(ground)
	rHas, rScore := m.rightHasPlayer(ground)
	bHas, bScore := m.backHasPlayer(ground)

	if fHas {
		positionSlice = append(positionSlice, position{
			position: FrontSide,
			score:    fScore,
		})
	}
	if lHas {
		positionSlice = append(positionSlice, position{
			position: LeftSide,
			score:    lScore,
		})
	}
	if rHas {
		positionSlice = append(positionSlice, position{
			position: RightSide,
			score:    rScore,
		})
	}
	if bHas {
		positionSlice = append(positionSlice, position{
			position: BackSide,
			score:    bScore,
		})
	}

	// high score
	var maxScorePosition int
	if len(positionSlice) == 0 {
		maxScorePosition = NoPlayer
	} else {
		sort.Slice(positionSlice, func(i, j int) bool {
			return positionSlice[i].score > positionSlice[j].score
		})
		maxScorePosition = positionSlice[0].position
	}
	m.enemyF = fHas
	m.enemyL = lHas
	m.enemyR = rHas
	m.enemyB = bHas
	m.highestScoreAroundMe = maxScorePosition

	//TODO Calculate the Highest Score on the ground
}

func (m *Myself) MyXY() (int, int) {
	return m.x, m.y
}

func (m *Myself) WasHit() bool {
	return m.wasHit
}

func (m *Myself) EnemyFront() bool {
	return m.enemyF
}

func (m *Myself) EnemyLeft() bool {
	return m.enemyL
}

func (m *Myself) EnemyRight() bool {
	return m.enemyR
}

func (m *Myself) EnemyBack() bool {
	return m.enemyB
}

func (m *Myself) HighestPlayerAroundMe() int {
	return m.highestScoreAroundMe
}

func (m *Myself) frontHasPlayer(ground *Ground) (bool, int) {
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

func (m *Myself) backHasPlayer(ground *Ground) (bool, int) {
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

func (m *Myself) leftHasPlayer(ground *Ground) (bool, int) {
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

func (m *Myself) rightHasPlayer(ground *Ground) (bool, int) {
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

func (m *Myself) MoveNorth(ground *Ground) string {
	switch m.direction {
	case NORTH:
		return m.MoveFront(ground)
	case SOUTH:
		return m.MoveRight(ground)
	case WEST:
		return m.MoveRight(ground)
	case EAST:
		return m.MoveLeft(ground)
	}
	return RIGHT
}

func (m *Myself) MoveSouth(ground *Ground) string {
	switch m.direction {
	case NORTH:
		return m.MoveRight(ground)
	case SOUTH:
		return m.MoveFront(ground)
	case WEST:
		return m.MoveLeft(ground)
	case EAST:
		return m.MoveRight(ground)
	}
	return RIGHT
}

func (m *Myself) MoveWest(ground *Ground) string {
	switch m.direction {
	case NORTH:
		return m.MoveLeft(ground)
	case SOUTH:
		return m.MoveRight(ground)
	case WEST:
		return m.MoveFront(ground)
	case EAST:
		return m.MoveRight(ground)
	}
	return RIGHT
}

func (m *Myself) MoveEast(ground *Ground) string {
	switch m.direction {
	case NORTH:
		return m.MoveRight(ground)
	case SOUTH:
		return m.MoveLeft(ground)
	case WEST:
		return m.MoveRight(ground)
	case EAST:
		return m.MoveFront(ground)
	}
	return RIGHT
}

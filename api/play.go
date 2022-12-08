package api

import (
	"gcp-hackathon-cloud-run/types"
)

type position struct {
	position int
	score    int
}

func Play(input types.ArenaUpdate) string {

	ground := types.InitGround(input.Arena.Dimensions, input.Arena.State)

	myselfData := input.Arena.State[input.Links.Self.Href]
	myself := types.InitMyself(myselfData.X, myselfData.Y, myselfData.Direction, myselfData.WasHit)
	myself.InitMyselfInGround(ground)
	return action(myself, ground)
}

func action(myself *types.Myself, ground *types.Ground) string {
	// Runaway
	if myself.WasHit() {
		if myself.EnemyBack() || myself.EnemyLeft() || myself.EnemyRight() {
			return myself.MoveFront(ground)
		}
	}

	// Attack
	if myself.EnemyFront() || myself.HighestPlayerAroundMe() == types.FrontSide {
		return types.THROW
	} else if myself.HighestPlayerAroundMe() == types.LeftSide {
		return myself.MoveLeft(ground)
	} else if myself.HighestPlayerAroundMe() == types.RightSide {
		return myself.MoveRight(ground)
	} else if myself.HighestPlayerAroundMe() == types.BackSide {
		return myself.MoveRight(ground)
	}

	return findPlayerNearBy(myself, ground)
}

// Find Player NearBy
func findPlayerNearBy(myself *types.Myself, ground *types.Ground) string {
	// find the highest score
	mx, my := myself.MyXY()
	hx, hy := ground.HighestPlayer()

	// find direction
	var direction string
	if hx > mx {
		direction = types.EAST
	} else if hx < mx {
		direction = types.WEST
	} else if hy > my {
		direction = types.SOUTH
	} else {
		direction = types.NORTH
	}

	switch direction {
	case types.NORTH:
		return myself.MoveNorth(ground)
	case types.SOUTH:
		return myself.MoveSouth(ground)
	case types.WEST:
		return myself.MoveWest(ground)
	case types.EAST:
		return myself.MoveEast(ground)
	}
	return myself.MoveNorth(ground)
}

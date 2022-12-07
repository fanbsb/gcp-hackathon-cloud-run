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
	if myself.HighestPlayerAroundMe() == types.FrontSide {
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
	return myself.MoveFront(ground)
}

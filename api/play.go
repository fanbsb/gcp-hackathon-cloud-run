package api

import (
	"gcp-hackathon-cloud-run/types"
	"sort"
)

type position struct {
	position int
	score    int
}

func Play(input types.ArenaUpdate) string {

	ground := types.InitGround(input.Arena.Dimensions, input.Arena.State)

	myselfData := input.Arena.State[input.Links.Self.Href]
	myself := types.InitMyself(myselfData.X, myselfData.Y, myselfData.Direction, myselfData.WasHit)
	return action(myself, ground)
}

func action(myself *types.Myself, ground *types.Ground) string {
	// Position
	var positionSlice []position
	fHas, fScore := myself.FrontHasPlayer(ground)
	lHas, lScore := myself.LeftHasPlayer(ground)
	rHas, rScore := myself.RightHasPlayer(ground)
	bHas, bScore := myself.BackHasPlayer(ground)

	if fHas {
		positionSlice = append(positionSlice, position{
			position: types.FrontSide,
			score:    fScore,
		})
	}
	if lHas {
		positionSlice = append(positionSlice, position{
			position: types.LeftSide,
			score:    lScore,
		})
	}
	if rHas {
		positionSlice = append(positionSlice, position{
			position: types.RightSide,
			score:    rScore,
		})
	}
	if bHas {
		positionSlice = append(positionSlice, position{
			position: types.BackSide,
			score:    bScore,
		})
	}

	// high score
	var maxScorePosition int
	if len(positionSlice) == 0 {
		maxScorePosition = types.NoPlayer
	} else {
		sort.Slice(positionSlice, func(i, j int) bool {
			return positionSlice[i].score > positionSlice[j].score
		})
		maxScorePosition = positionSlice[0].position
	}

	// Attack first
	switch maxScorePosition {
	case types.FrontSide:
		return types.THROW
	case types.LeftSide:
		return types.LEFT
	case types.RightSide:
		return types.RIGHT
	case types.BackSide:
		return types.LEFT
	case types.NoPlayer:
		//TODO Find Player
	}

	// Runaway
	//if myself.WasHit() {
	//	return wasHit(myself, ground)
	//}

	return findPlayerNearBy(myself, ground)
}

// Find the player and runaway
func wasHit(myself *types.Myself, ground *types.Ground) string {
	if myself.CanMoveFront(ground) {
		return types.FORWARD
	} else if myself.CanMoveLeft(ground) {
		return types.LEFT
	} else if myself.CanMoveRight(ground) {
		return types.RIGHT
	} else {
		return types.LEFT
	}
}

// Find Player NearBy
func findPlayerNearBy(myself *types.Myself, ground *types.Ground) string {
	if myself.CanMoveFront(ground) {
		return types.FORWARD
	} else if myself.CanMoveLeft(ground) {
		return types.LEFT
	} else if myself.CanMoveRight(ground) {
		return types.RIGHT
	}
	return types.RIGHT
}

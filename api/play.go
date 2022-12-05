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

	ground.Print()

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
	if maxScorePosition == types.FrontSide {
		return types.THROW
	}

	// Runaway
	if myself.WasHit() {
		switch maxScorePosition {
		case types.LeftSide:
			return myself.MoveFront(ground)
		case types.RightSide:
			return myself.MoveFront(ground)
		case types.BackSide:
			return myself.MoveFront(ground)
		case types.FrontSide:
			// Already attack
		case types.NoPlayer:
			// Impossible
		}
	}

	return findPlayerNearBy(myself, ground)
}

// Find Player NearBy
func findPlayerNearBy(myself *types.Myself, ground *types.Ground) string {
	return myself.MoveFront(ground)
}

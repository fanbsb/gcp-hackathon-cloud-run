package api

import (
	"gcp-hackathon-cloud-run/types"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	positionSlice := []position{{
		position: types.LeftSide,
		score:    100,
	}, {
		position: types.RightSide,
		score:    200,
	}, {
		position: types.FrontSide,
		score:    2,
	}, {
		position: types.BackSide,
		score:    3,
	}}

	sort.Slice(positionSlice, func(i, j int) bool {
		return positionSlice[i].score > positionSlice[j].score
	})

	if positionSlice[0].position != types.RightSide {
		t.Fail()
	} else {
		t.Log("Success")
	}
}

package types

const (
	// Direction
	NORTH = "N"
	SOUTH = "S"
	WEST  = "W"
	EAST  = "E"

	// Player Position
	FrontSide = iota
	LeftSide
	RightSide
	BackSide
	NoPlayer

	// Command
	FORWARD = "F"
	LEFT    = "L"
	RIGHT   = "R"
	THROW   = "T"
)

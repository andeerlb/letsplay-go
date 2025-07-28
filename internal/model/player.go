package model

type GameType string

const (
	Fut11  GameType = "fut11"
	Fut7   GameType = "fut7"
	Futsal GameType = "futsal"
)

type PlayerPosition string

var positionsFut11 = []PlayerPosition{
	"goalkeeper",
	"center-back",
	"right-back",
	"left-back",
	"defensive-midfielder",
	"central-midfielder",
	"attacking-midfielder",
	"right-winger",
	"left-winger",
	"striker",
	"second-striker",
}

var positionsFut7 = []PlayerPosition{
	"goalkeeper",
	"fixo",
	"right-wing",
	"left-wing",
	"midfielder",
	"pivot",
}

var positionsFutsal = []PlayerPosition{
	"goalkeeper",
	"fixo",
	"right-wing",
	"left-wing",
	"midfielder",
	"pivot",
}

type GameInfo struct {
	Type     GameType       `json:"type"`
	Position PlayerPosition `json:"position"`
}

package entity

import "time"

type Game struct {
	ID          uint
	CategoryID  uint
	QuestionIDs []uint
	Players     []Player
	StartTime   time.Time
}

type Player struct {
	ID      uint
	UserID  uint
	GameID  uint
	Score   uint
	answers []PlayerAnswer
}

type PlayerAnswer struct {
	ID         uint
	PlayerID   uint
	QuestionID uint
	Choice     PossibleAnswer
}

package entity

type Question struct {
	ID              uint
	Text            string
	PossibleAnswers []PossibleAnswer
	CurrentAnswerID uint
	Difficulty      QuestionDifficulty
	CategoryID      uint
}

type PossibleAnswer struct {
	ID     uint
	Text   string
	Choice PossibleAnswerChoice
}

type PossibleAnswerChoice uint8

func (p PossibleAnswerChoice) IsValid() bool {
	return p >= PossibleAnswerA && p <= PossibleAnswerD
}

const (
	PossibleAnswerA PossibleAnswerChoice = iota + 1
	PossibleAnswerB
	PossibleAnswerC
	PossibleAnswerD
)

type QuestionDifficulty uint8

func (q QuestionDifficulty) IsValid() bool {
	return q >= QuestionDifficultyEasy && q <= QuestionDifficultyHard
}

const (
	QuestionDifficultyEasy QuestionDifficulty = iota + 1
	QuestionDifficultyMedium
	QuestionDifficultyHard
)

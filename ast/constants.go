package ast

type PuzzleType string

const Grid PuzzleType = "grid"

type ClueType string

const (
	Columns ClueType = "columns"
	Rows    ClueType = "rows"
)

type SolutionType string

const Goal SolutionType = "goal"

package webpbn

type PuzzleType string

const PuzzleGrid PuzzleType = "grid"

type ClueType string

const (
	ClueColumns ClueType = "columns"
	ClueRows    ClueType = "rows"
)

type SolutionType string

const (
	SolutionGoal     SolutionType = "goal"
	SolutionSolution SolutionType = "solution"
	SolutionSaved    SolutionType = "saved"
)

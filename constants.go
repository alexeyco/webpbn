package webpbn

type PuzzleType string

const (
	PuzzleGrid     PuzzleType = "grid"
	PuzzleTriddler PuzzleType = "triddler"
)

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

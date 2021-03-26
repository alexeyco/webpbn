package ast

// PuzzleType declares a type of the puzzle.
type PuzzleType string

// Grid used for a puzzle where cells are square and there will be a set of row clues and a set of column clue.
const Grid PuzzleType = "grid"

// ClueType declares a type of clue.
type ClueType string

const (
	// Columns used for a clue that stores columns.
	Columns ClueType = "columns"

	// Rows used for a clue that stores rows.
	Rows ClueType = "rows"
)

// SolutionType declares a type of the solution.
type SolutionType string

// Goal used for the goal solution intended by the designer of the puzzle.
const Goal SolutionType = "goal"

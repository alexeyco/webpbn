package validator

import "github.com/alexeyco/webpbn/ast"

// Validator puzzle validator.
type Validator struct{}

// Validate validates a set of puzzles.
func (v *Validator) Validate(puzzleSet *ast.PuzzleSet) error {
	return nil
}

// New returns new validator instance.
func New() *Validator {
	return &Validator{}
}

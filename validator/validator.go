package validator

import (
	"github.com/alexeyco/webpbn/ast"
)

type Validator struct {
}

func (v *Validator) Validate(puzzleSet *ast.PuzzleSet) error {
	return nil
}

func New() *Validator {
	return &Validator{}
}

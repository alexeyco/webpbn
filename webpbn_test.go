package webpbn_test

import (
	"testing"

	"github.com/alexeyco/webpbn/ast"
)

type validatorMock struct {
	onValidate func(puzzleSet *ast.PuzzleSet) error
}

func (v *validatorMock) Validate(puzzleSet *ast.PuzzleSet) error {
	if v.onValidate != nil {
		return v.onValidate(puzzleSet)
	}

	return nil
}

func TestParse(t *testing.T) {
	t.Parallel()

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()
	})

	t.Run("IncorrectXML", func(t *testing.T) {
		t.Parallel()
	})
}

func TestParseFile(t *testing.T) {
	t.Parallel()

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()
	})

	t.Run("FileError", func(t *testing.T) {
		t.Parallel()
	})
}

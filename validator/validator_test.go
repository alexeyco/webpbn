package validator_test

import (
	"reflect"
	"testing"

	"github.com/alexeyco/webpbn/ast"
	"github.com/alexeyco/webpbn/validator"
)

var testData = [...]struct {
	name      string
	puzzleSet *ast.PuzzleSet
	err       error
}{}

func TestValidator_Validate(t *testing.T) {
	t.Parallel()

	v := validator.New()

	for _, testDatum := range testData {
		testDatum := testDatum

		t.Run(testDatum.name, func(t *testing.T) {
			t.Parallel()

			err := v.Validate(testDatum.puzzleSet)

			if !reflect.DeepEqual(err, testDatum.err) {
				t.Error(`Errors should be equal`)
			}
		})
	}
}

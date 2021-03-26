package webpbn_test

import (
	"reflect"
	"testing"

	"github.com/alexeyco/webpbn"
)

func TestWithValidator(t *testing.T) {
	t.Parallel()

	options := webpbn.Options{}
	validator := validatorMock{}

	webpbn.WithValidator(&validator)(&options)

	if !reflect.DeepEqual(options.Validator, &validator) {
		t.Error(`Validators should be equal`)
	}
}

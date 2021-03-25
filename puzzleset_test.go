package webpbn_test

import (
	"testing"

	"github.com/alexeyco/webpbn"
)

var colors = webpbn.Colors{
	{
		Name: "white",
		Char: '.',
		Hex:  "fff",
	},
	{
		Name: "black",
		Char: 'x',
		Hex:  "000",
	},
}

func TestColors_GetByName(t *testing.T) {
	t.Parallel()

	t.Run("Exists", func(t *testing.T) {
		t.Parallel()

		color, ok := colors.GetByName("white")
		if !ok {
			t.Error(`Should be true`)
		}

		if color.Name != "white" {
			t.Errorf(`Should be "white", "%s" given`, color.Name)
		}

		if color.Char != '.' {
			t.Errorf(`Should be ".", "%s" given`, string(color.Char))
		}

		if color.Hex != "fff" {
			t.Errorf(`Should be "fff", "%s" given`, color.Hex)
		}
	})

	t.Run("DoesNotExist", func(t *testing.T) {
		t.Parallel()

		color, ok := colors.GetByName("red")
		if ok {
			t.Error(`Should be false`)
		}

		if color != nil {
			t.Error(`Should be nil`)
		}
	})
}

func TestColors_GetByChar(t *testing.T) {
	t.Parallel()

	t.Run("Exists", func(t *testing.T) {
		t.Parallel()

		color, ok := colors.GetByChar('x')
		if !ok {
			t.Error(`Should be true`)
		}

		if color.Name != "black" {
			t.Errorf(`Should be "black", "%s" given`, color.Name)
		}

		if color.Char != 'x' {
			t.Errorf(`Should be "x", "%s" given`, string(color.Char))
		}

		if color.Hex != "000" {
			t.Errorf(`Should be "000", "%s" given`, color.Hex)
		}
	})

	t.Run("DoesNotExist", func(t *testing.T) {
		t.Parallel()

		color, ok := colors.GetByChar('-')
		if ok {
			t.Error(`Should be false`)
		}

		if color != nil {
			t.Error(`Should be nil`)
		}
	})
}

func TestClues_GetByType(t *testing.T) {
	t.Parallel()

	t.Run("Exists", func(t *testing.T) {
		t.Parallel()

	})

	t.Run("DoesNotExist", func(t *testing.T) {
		t.Parallel()

	})
}

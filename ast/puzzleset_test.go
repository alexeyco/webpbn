package ast_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/alexeyco/webpbn/ast"
)

var colors = ast.Colors{
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

func TestChar_MarshalText(t *testing.T) {
	t.Parallel()

	ch := ast.Char('x')
	actual, err := ch.MarshalText()

	if err != nil {
		t.Errorf(`Error should be nil, "%v" given`, err)
	}

	if string(actual) != "x" {
		t.Errorf(`Should be "x", "%s" given`, string(actual))
	}
}

func TestChar_UnmarshalText(t *testing.T) {
	t.Parallel()

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		var ch ast.Char
		if err := ch.UnmarshalText([]byte("x")); err != nil {
			t.Errorf(`Error should be nil, "%v" given`, err)
		}

		if rune(ch) != 'x' {
			t.Errorf(`Char should be 'x', '%s' given`, string(ch))
		}
	})

	t.Run("ErrorCauseEmpty", func(t *testing.T) {
		t.Parallel()

		var ch ast.Char
		err := ch.UnmarshalText([]byte{})

		if err == nil {
			t.Error(`Error should not be nil`)
		}

		if !errors.Is(err, ast.ErrUnmarshal) {
			t.Errorf(`Error should be ast.ErrUnmarshal, "%v" given`, err)
		}

		if ch != 0 {
			t.Errorf(`Char should be blank, '%s' given`, string(ch))
		}
	})

	t.Run("ErrorCauseTooManyCharacters", func(t *testing.T) {
		t.Parallel()

		var ch ast.Char
		err := ch.UnmarshalText([]byte("xx"))

		if err == nil {
			t.Error(`Error should not be nil`)
		}

		if !errors.Is(err, ast.ErrUnmarshal) {
			t.Errorf(`Error should be ast.ErrUnmarshal, "%v" given`, err)
		}

		if ch != 0 {
			t.Errorf(`Char should be blank, '%s' given`, string(ch))
		}
	})
}

func TestClues_GetByType(t *testing.T) {
	t.Parallel()

	clues := ast.Clues{
		{
			Type: "foo",
		},
	}

	t.Run("Exists", func(t *testing.T) {
		t.Parallel()

		c, ok := clues.GetByType("foo")

		if !ok {
			t.Error(`Should be true`)
		}

		if !reflect.DeepEqual(c, &clues[0]) {
			t.Error(`Should be equal to existing clue`)
		}
	})

	t.Run("DoesNotExist", func(t *testing.T) {
		t.Parallel()

		c, ok := clues.GetByType("bar")

		if ok {
			t.Error(`Should be false`)
		}

		if c != nil {
			t.Error(`Should be nil`)
		}
	})
}

func TestImage_MarshalText(t *testing.T) {
	t.Parallel()

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		image := ast.Image{
			{'a', 'b', 'c'},
			{'d', 'e', 'f'},
			{'g', 'h', 'i'},
			{'j', 'k', 'l'},
		}

		expected := "|abc||def||ghi||jkl|"

		b, err := image.MarshalText()

		if err != nil {
			t.Errorf(`Should be nil, "%v" given`, err)
		}

		if string(b) != expected {
			t.Errorf(`Should be "%s", "%s" given`, expected, string(b))
		}
	})

	t.Run("Empty", func(t *testing.T) {
		var image ast.Image
		b, err := image.MarshalText()

		if err != nil {
			t.Errorf(`Should be nil, "%v" given`, err)
		}

		if len(b) > 0 {
			t.Errorf(`Should be blank, "%s" given`, string(b))
		}
	})
}

func TestImage_UnmarshalText(t *testing.T) {
	t.Parallel()

	expected := ast.Image{
		{'a', 'b', 'c'},
		{'d', 'e', 'f'},
		{'g', 'h', 'i'},
		{'j', 'k', 'l'},
	}

	t.Run("Ok", func(t *testing.T) {
		t.Parallel()

		var image ast.Image
		if err := image.UnmarshalText([]byte("|abc||def||ghi||jkl|")); err != nil {
			t.Errorf(`Should be nil, "%v" given`, err)
		}

		if !reflect.DeepEqual(image, expected) {
			t.Log(image)
			t.Error(`Should be equal`)
		}
	})

	t.Run("OkWithLineBreaks", func(t *testing.T) {
		t.Parallel()

		var image ast.Image
		if err := image.UnmarshalText([]byte("\n|abc|\n|def|\n|ghi|\n|jkl|\n")); err != nil {
			t.Errorf(`Should be nil, "%v" given`, err)
		}

		if !reflect.DeepEqual(image, expected) {
			t.Error(`Should be equal`)
		}
	})
}

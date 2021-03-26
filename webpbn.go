package webpbn

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/alexeyco/webpbn/ast"
)

// Validator puzzle validator interface.
type Validator interface {
	Validate(puzzleSet *ast.PuzzleSet) error
}

// Parse puzzle set from reader.
func Parse(r io.Reader, options ...Option) (*ast.PuzzleSet, error) {
	var puzzleSet ast.PuzzleSet
	if err := xml.NewDecoder(r).Decode(&puzzleSet); err != nil {
		return nil, fmt.Errorf(`can't decode puzzle file: %w`, err)
	}

	return &puzzleSet, nil
}

// ParseFile loads puzzle set from file by name.
func ParseFile(name string, options ...Option) (*ast.PuzzleSet, error) {
	// nolint:gosec
	file, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf(`can't read puzzle file: %w`, err)
	}

	defer func() {
		_ = file.Close()
	}()

	return Parse(file, options...)
}

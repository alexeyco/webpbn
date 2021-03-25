package webpbn

import (
	"encoding/xml"
	"io"
	"os"
)

// Parse puzzleset from reader.
func Parse(r io.Reader) (*PuzzleSet, error) {
	var puzzleSet PuzzleSet
	if err := xml.NewDecoder(r).Decode(&puzzleSet); err != nil {
		return nil, err
	}

	return &puzzleSet, nil
}

// ParseFile loads puzzleset from file by name.
func ParseFile(name string) (*PuzzleSet, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = file.Close()
	}()

	return Parse(file)
}

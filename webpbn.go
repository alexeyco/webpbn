package webpbn

import (
	"io"
	"os"
)

func Parse(r io.Reader) (*Doc, error) {
	return nil, nil
}

func ParseFile(name string) (*Doc, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = file.Close()
	}()

	return Parse(file)
}

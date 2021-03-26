package ast

import (
	"encoding/xml"
	"fmt"
	"strings"
	"unicode/utf8"
)

// PuzzleSet the root structure of the document. Includes all puzzles.
type PuzzleSet struct {
	XMLName xml.Name `xml:"puzzleset"`
	Puzzles []Puzzle `xml:"puzzle"`
}

// Puzzle a puzzle in the set of puzzles.
type Puzzle struct {
	Type            PuzzleType `xml:"type,attr"`
	DefaultColor    string     `xml:"defaultcolor,attr,omitempty"`
	BackgroundColor string     `xml:"backgroundcolor,attr,omitempty"`
	Source          string     `xml:"source,omitempty"`
	ID              string     `xml:"id,omitempty"`
	Author          string     `xml:"author,omitempty"`
	AuthorID        string     `xml:"authorid,omitempty"`
	Copyright       string     `xml:"copyright,omitempty"`
	Description     string     `xml:"description,omitempty"`
	Colors          Colors     `xml:"color"`
	Clues           Clues      `xml:"clues"`
	Solution        *Solution  `xml:"solution,omitempty"`
}

// Colors collection of colors.
type Colors []Color

// GetByName returns a color from the collection by his name.
func (c Colors) GetByName(name string) (*Color, bool) {
	for _, color := range c {
		if color.Name == name {
			return &color, true
		}
	}

	return nil, false
}

// GetByChar returns a color from the collection by his character.
func (c Colors) GetByChar(char Char) (*Color, bool) {
	for _, color := range c {
		if color.Char == char {
			return &color, true
		}
	}

	return nil, false
}

// Color defines a color name used in the puzzle.
type Color struct {
	Name string `xml:"name,attr"`
	Char Char   `xml:"char,attr"`
	Hex  string `xml:",chardata"`
}

// Char defines a color unique character.
type Char rune

// MarshalText encodes the character into UTF-8-encoded text and returns the result.
func (c Char) MarshalText() ([]byte, error) {
	return []byte(string(c)), nil
}

// UnmarshalText decodes the character from UTF-8-encoded text.
func (c *Char) UnmarshalText(b []byte) error {
	cnt := utf8.RuneCount(b)
	if cnt == 0 {
		return fmt.Errorf(`%w empty string: should be a single char`, ErrUnmarshal)
	}

	if cnt > 1 {
		return fmt.Errorf(`%w "%s": should be a single char`, ErrUnmarshal, string(b))
	}

	r, _ := utf8.DecodeRune(b)
	*c = Char(r)

	return nil
}

// Clues collection of clues.
type Clues []Clue

// GetByType returns a clue by his type.
func (c Clues) GetByType(t ClueType) (*Clue, bool) {
	for _, clue := range c {
		if clue.Type == t {
			return &clue, true
		}
	}

	return nil, false
}

// Clue defines a clue used in the puzzle.
type Clue struct {
	Type  ClueType `xml:"type,attr"`
	Lines []Line   `xml:"line"`
}

// Line defines clue line.
type Line struct {
	Counts []Count `xml:"count"`
}

// Count defines line counts.
type Count struct {
	Count int    `xml:",chardata"`
	Color string `xml:"color,attr,omitempty"`
}

// Solution defines puzzle solution.
type Solution struct {
	Type  SolutionType `xml:"type,attr"`
	Image Image        `xml:"image"`
}

// Image defines solution image.
type Image [][]Char

// MarshalText encodes the image into UTF-8-encoded text and returns the result.
func (i Image) MarshalText() (b []byte, err error) {
	if len(i) == 0 {
		return
	}

	lines := make([]string, len(i))

	for n, chars := range i {
		line := make([]rune, len(chars))
		for x, r := range chars {
			line[x] = rune(r)
		}

		lines[n] = fmt.Sprintf("|%s|", string(line))
	}

	return []byte(strings.Join(lines, "")), nil
}

// UnmarshalText decodes the image from UTF-8-encoded text.
func (i *Image) UnmarshalText(text []byte) error {
	txt := strings.TrimSpace(string(text))
	txt = strings.ReplaceAll(txt, "||", "|\n|")

	lines := strings.Split(txt, "\n")
	img := make([][]Char, len(lines))

	for n, line := range lines {
		line = strings.Trim(line, "|")
		runes := []rune(line)

		img[n] = make([]Char, len(runes))
		for m, r := range runes {
			img[n][m] = Char(r)
		}
	}

	*i = img

	return nil
}

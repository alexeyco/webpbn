package ast

import (
	"encoding/xml"
	"fmt"
	"strings"
	"unicode/utf8"
)

type PuzzleSet struct {
	XMLName xml.Name `xml:"puzzleset"`
	Puzzles []Puzzle `xml:"puzzle"`
}

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

type Colors []Color

func (c Colors) GetByName(name string) (*Color, bool) {
	for _, color := range c {
		if color.Name == name {
			return &color, true
		}
	}

	return nil, false
}

func (c Colors) GetByChar(char Char) (*Color, bool) {
	for _, color := range c {
		if color.Char == char {
			return &color, true
		}
	}

	return nil, false
}

type Color struct {
	Name string `xml:"name,attr"`
	Char Char   `xml:"char,attr"`
	Hex  string `xml:",chardata"`
}

type Char rune

func (c Char) MarshalText() ([]byte, error) {
	return []byte(string(c)), nil
}

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

type Clues []Clue

func (c Clues) GetByType(t ClueType) (*Clue, bool) {
	for _, clue := range c {
		if clue.Type == t {
			return &clue, true
		}
	}

	return nil, false
}

type Clue struct {
	Type  ClueType `xml:"type,attr"`
	Lines []Line   `xml:"line"`
}

type Line struct {
	Counts []Count `xml:"count"`
}

type Count struct {
	Count int    `xml:",chardata"`
	Color string `xml:"color,attr,omitempty"`
}

type Solution struct {
	Type  SolutionType `xml:"type,attr"`
	Image Image        `xml:"image"`
}

type Image [][]Char

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

func (i *Image) UnmarshalText(text []byte) error {
	txt := strings.TrimSpace(string(text))
	txt = strings.Replace(txt, "||", "|\n|", -1)

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

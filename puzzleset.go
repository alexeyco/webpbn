package webpbn

import "encoding/xml"

type PuzzleSet struct {
	XMLName xml.Name `xml:"puzzleset"`
	Puzzles []Puzzle `xml:"puzzle"`
}

func (s PuzzleSet) Validate() error {
	for _, p := range s.Puzzles {
		if err := p.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type Puzzle struct {
	Type            PuzzleType `xml:"type,attr"`
	DefaultColor    string     `xml:"defaultcolor,attr"`
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

func (p Puzzle) Size() (int, int) {
	var (
		columns int
		rows    int
	)

	if clue, ok := p.Clues.GetByType(ClueColumns); ok {
		columns = len(clue.Lines)
	}

	if clue, ok := p.Clues.GetByType(ClueRows); ok {
		rows = len(clue.Lines)
	}

	return columns, rows
}

func (p Puzzle) Validate() error {
	return nil
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

func (c Colors) GetByChar(char string) (*Color, bool) {
	for _, color := range c {
		if color.Char == char {
			return &color, true
		}
	}

	return nil, false
}

type Color struct {
	Name string `xml:"name,attr"`
	Char string `xml:"char,attr"`
	Hex  string `xml:",chardata"`
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
	Type  SolutionType `xml:"type"`
	Image Image        `xml:"image"`
}

type Image struct {
}

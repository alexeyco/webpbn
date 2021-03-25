package webpbn

import "encoding/xml"

type PuzzleSet struct {
	XMLName xml.Name `xml:"puzzleset"`
	Puzzles []Puzzle `xml:"puzzle"`
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
	Colors          []Color    `xml:"color"`
	Clues           []Clue     `xml:"clues"`
	Solution        Solution   `xml:"solution"`
}

type Color struct {
	Name string `xml:"name,attr"`
	Char string `xml:"char,attr"`
	Hex  string `xml:",chardata"`
}

type Clue struct {
	Type  ClueType `xml:"type,attr"`
	Lines []Line   `xml:"line"`
}

type Line struct {
	Counts []Count `xml:"count"`
}

type Count struct {
	Count int `xml:",chardata"`
}

type Solution struct {
	Type  SolutionType `xml:"type"`
	Image Image        `xml:"image"`
}

type Image struct {
}

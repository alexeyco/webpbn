package webpbn

type Doc struct {
	Puzzleset Puzzleset `xml:"puzzleset"`
}

type Puzzleset struct {
	Puzzles []Puzzle `xml:"puzzle"`
}

type Puzzle struct {
	Type            PuzzleType `xml:"type"`
	DefaultColor    string     `xml:"defaultcolor"`
	BackgroundColor string     `xml:"backgroundcolor"`
	Source          string     `xml:"source"`
	ID              string     `xml:"id"`
	Author          string     `xml:"author"`
	AuthorID        string     `xml:"authorid"`
	Copyright       string     `xml:"copyright"`
	Description     string     `xml:"description"`
	Colors          []Color    `xml:"color"`
	Clues           []Clue     `xml:"clues"`
	Solution        Solution   `xml:"solution"`
}

type Color struct {
	Name string `xml:"name"`
	Char string `xml:"char"`
	Hex  string `xml:"hex"`
}

type Clue struct {
	Type  ClueType `xml:"type"`
	Lines []Line   `xml:"line"`
}

type Line struct {
	Counts []int `xml:"count"`
}

type Solution struct {
	Type  SolutionType `xml:"type"`
	Image Image        `xml:"image"`
}

type Image struct {
}

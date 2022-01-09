package dto

type PDFDoc struct {
	Orientation string
	Unit        string
	Size        string
	FontDir     string
}

type TitleStruct struct {
	Title string
	Font  FontStruct
	Cell  CellStruct
}

type FontStruct struct {
	Family string
	Style  string
	Size   float64
}

type CellStruct struct {
	Width   float64
	Height  float64
	Content string
	Border  string
	Ln      int
	Align   string
	Fill    bool
	Link    int
	LinkStr string
}

type ColorStruct struct {
	R int
	G int
	B int
}

type TableStruct struct {
	HeaderFont FontStruct
	BodyFont   FontStruct
	Color      ColorStruct
}

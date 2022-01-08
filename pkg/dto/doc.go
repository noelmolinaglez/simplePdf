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
}

type FontStruct struct {
	Family string
	Style  string
	Size   int
}

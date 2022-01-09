package helpers

import "github.com/noelmolinaglez/simplePdf/pkg/dto"

func CreateSimpleCell(content string) dto.CellStruct {
	return dto.CellStruct{
		Width:   40,
		Height:  10,
		Content: content,
		Border:  "1",
		Ln:      0,
		Align:   "C",
		Fill:    false,
		Link:    0,
		LinkStr: "",
	}

}

func CreateSimpleTitle(content string) dto.TitleStruct {
	return dto.TitleStruct{
		Title: content,
		Font: dto.FontStruct{
			Family: "Times",
			Style:  "B",
			Size:   22,
		},
		Cell: dto.CellStruct{
			Width:   0,
			Height:  10,
			Content: content,
			Border:  "1",
			Ln:      0,
			Align:   "C",
			Fill:    false,
			Link:    0,
			LinkStr: "",
		},
	}

}

func CreateSimpleDoc() dto.PDFDoc {
	return dto.PDFDoc{
		Orientation: "L",
		Unit:        "mm",
		Size:        "letter",
		FontDir:     "",
	}
}

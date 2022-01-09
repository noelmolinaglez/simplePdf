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

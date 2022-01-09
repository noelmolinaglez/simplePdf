package interfaces

import "github.com/noelmolinaglez/simplePdf/pkg/dto"

type SimpleDoc interface {
	GetHeaders() []dto.CellStruct
	GetValues() []dto.CellStruct
}

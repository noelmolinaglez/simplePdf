package templates

import (
	"github.com/noelmolinaglez/simplePdf/pkg/dto"
	"github.com/noelmolinaglez/simplePdf/pkg/helpers"
	"github.com/noelmolinaglez/simplePdf/pkg/interfaces"
)

func GenerateSimpleTable(doc dto.PDFDoc, title dto.TitleStruct, data []interfaces.SimpleDoc, table dto.TableStruct) {

	pdf := helpers.CreateDoc(doc, title)

	pdf = helpers.GenericTable(pdf, data, table)

	_ = helpers.SaveDoc(pdf, title.Title)
}

package doc

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"github.com/noelmolinaglez/simplePdf/pkg/dto"
)

//Creates a one page pdf with a title
func CreateDoc(doc dto.PDFDoc, title dto.TitleStruct) *gofpdf.Fpdf {
	pdf := gofpdf.New(doc.Orientation, doc.Unit, doc.Size, doc.FontDir)

	pdf.SetTitle(title.Title, true)

	pdf.AddPage()

	pdf.SetFont(title.Font.Family, title.Font.Style, title.Font.Size)
	pdf.Cell(title.Cell.Width, title.Cell.Height, title.Title)

	pdf.Ln(12)

	return pdf
}

func SaveDoc(pdf *gofpdf.Fpdf, fileName string) error {
	name := fmt.Sprintf("%s.pdf", fileName)
	return pdf.OutputFileAndClose(name)
}

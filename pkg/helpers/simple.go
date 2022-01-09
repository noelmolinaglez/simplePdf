package helpers

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
	"github.com/noelmolinaglez/simplePdf/pkg/dto"
	"github.com/noelmolinaglez/simplePdf/pkg/interfaces"
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

func GenericTable(pdf *gofpdf.Fpdf, data []interfaces.SimpleDoc, table dto.TableStruct) *gofpdf.Fpdf {
	pdf.SetFont(table.Font.Family, table.Font.Style, table.Font.Size)
	pdf.SetFillColor(table.Color.R, table.Color.G, table.Color.B)

	headers := data[0].GetHeaders()
	size := GetRowWidth(pdf)

	for _, value := range headers {
		pdf.CellFormat(value.Width*size, value.Height, value.Content, value.Border, value.Ln, value.Align, value.Fill, value.Link, value.LinkStr)
	}
	pdf.Ln(-1)
	for _, record := range data {
		for _, value := range record.GetValues() {
			pdf.CellFormat(value.Width*size, value.Height, value.Content, value.Border, value.Ln, value.Align, value.Fill, value.Link, value.LinkStr)
		}
		pdf.Ln(-1)
	}

	pdf.Ln(12)

	return pdf
}

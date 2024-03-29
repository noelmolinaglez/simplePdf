package helpers

import (
	"fmt"
	"github.com/jung-kurt/gofpdf"
)

func GetRowWidth(pdf *gofpdf.Fpdf) float64 {
	width, _ := pdf.GetPageSize()
	left, _, right, _ := pdf.GetMargins()
	width = width - (left + right)
	return width / 12
}

func DrawCurrentPageNumber(pdf *gofpdf.Fpdf) {
	pdf.SetY(-15)
	pdf.SetFont("Arial", "I", 8)
	pdf.CellFormat(0, 10, fmt.Sprintf(" %d/{nb}", pdf.PageNo()),
		"", 0, "C", false, 0, "")
}

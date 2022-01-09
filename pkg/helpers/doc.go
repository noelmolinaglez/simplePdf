package helpers

import "github.com/jung-kurt/gofpdf"

func GetRowWidth(pdf *gofpdf.Fpdf) float64 {
	width, _ := pdf.GetPageSize()
	left, _, right, _ := pdf.GetMargins()
	width = width - (left + right)
	return width / 12
}

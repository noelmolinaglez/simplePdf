package helpers

import "github.com/jung-kurt/gofpdf"

func GetRowWidth(pdf *gofpdf.Fpdf) float64 {
	width, _ := pdf.GetPageSize()
	return width / 12
}

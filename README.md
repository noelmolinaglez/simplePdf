# SimplePdf
A simple wrapper to the gofpdf package to help us create beautiful pdfs documents in an easy way.

Example 1 Creating a table report for the departments list :

* Implements the SimpleDoc interface

```
type SimpleDoc interface {
	GetHeaders() []dto.CellStruct
	GetValues() []dto.CellStruct
}

type Department struct {
	DepartmentID uint   
	Name         string 
	GroupName    string 
	ModifiedDate string 
}

func (Department) GetHeaders() []dto.CellStruct {

	headers := make([]dto.CellStruct, 4)

	headers[0] = dto.CellStruct{
		Width:   2,
		Height:  10,
		Content: "DepartmentId",
		Border:  "1",
		Ln:      0,
		Align:   "C",
		Fill:    true,
		Link:    0,
		LinkStr: "",
	}
	headers[1] = dto.CellStruct{
		Width:   3,
		Height:  10,
		Content: "Name",
		Border:  "1",
		Ln:      0,
		Align:   "C",
		Fill:    true,
		Link:    0,
		LinkStr: "",
	}
	headers[2] = dto.CellStruct{
		Width:   4,
		Height:  10,
		Content: "GroupName",
		Border:  "1",
		Ln:      0,
		Align:   "C",
		Fill:    true,
		Link:    0,
		LinkStr: "",
	}
	headers[3] = dto.CellStruct{
		Width:   3,
		Height:  10,
		Content: "ModifiedDate",
		Border:  "1",
		Ln:      0,
		Align:   "C",
		Fill:    true,
		Link:    0,
		LinkStr: "",
	}

	return headers
}

func (d Department) GetValues() []dto.CellStruct {

	values := make([]dto.CellStruct, 4)
	values[0] = dto.CellStruct{
		Width:   2,
		Height:  10,
		Content: strconv.Itoa(int(d.DepartmentID)),
		Border:  "1",
		Ln:      0,
		Align:   "C",
		Fill:    false,
		Link:    0,
		LinkStr: "",
	}

	values[1] = dto.CellStruct{
		Width:   3,
		Height:  10,
		Content: d.Name,
		Border:  "1",
		Ln:      0,
		Align:   "C",
		Fill:    false,
		Link:    0,
		LinkStr: "",
	}

	values[2] = dto.CellStruct{
		Width:   4,
		Height:  10,
		Content: d.GroupName,
		Border:  "1",
		Ln:      0,
		Align:   "C",
		Fill:    false,
		Link:    0,
		LinkStr: "",
	}

	values[3] = dto.CellStruct{
		Width:   3,
		Height:  10,
		Content: d.ModifiedDate,
		Border:  "1",
		Ln:      0,
		Align:   "C",
		Fill:    false,
		Link:    0,
		LinkStr: "",
	}

	return values
}

```

* Creating the docs
```
	doc := helpers.CreateSimpleDoc()

```
* Creating the title
```
	title := helpers.CreateSimpleTitle("Sample")

```
* Creating and formatting the table

```
table := dto.TableStruct{
		BodyFont: dto.FontStruct{
			Family: "Times",
			Style:  "",
			Size:   12,
		},
		BodyColor: dto.ColorStruct{
			R: 255,
			G: 255,
			B: 255,
		},
		HeaderFont: dto.FontStruct{
			Family: "Times",
			Style:  "B",
			Size:   16,
		},
		HeaderColor: dto.ColorStruct{
			R: 240,
			G: 240,
			B: 240,
		},
	}
  ```
  * Converting the departments list to SimpleDoc list
  
  ```
  var departments [] Department
  data := make([]interfaces.SimpleDoc, len(departments))
	for i, k := range results {
		data[i] = k
	}
  ```
  * Saving the doc
  ```
  templates.GenerateSimpleTable(doc, title, data, table)

  ```
  Output:
  ![image](https://user-images.githubusercontent.com/43006548/148706148-2acbcd95-d0c6-4261-9fb4-8192438bfcee.png)


package main

import (
	//"bufio"
	//"bytes"

	"io/ioutil"

	"github.com/signintech/gopdf"
)

func main() {
    pdf := gopdf.GoPdf{}
    pdf.Start(gopdf.Config{PageSize: gopdf.Rect{W: 595, H: 420}})

    pdf.AddPage()

    err := pdf.AddTTFFont("font", "font.ttf")
    if err != nil {
        panic(err)
    }

    err = pdf.SetFont("font", "", 15)
    if err != nil {
        panic(err)
    }

    // Import page 1
    tpl1 := pdf.ImportPage("test.pdf", 1, "/MediaBox")

    // Draw pdf onto page
    pdf.UseImportedTemplate(tpl1, 0, 0, 595, 420)
    
    pdf.SetXY(55, 90)
    pdf.Cell(nil, "1234567890123")

    pdf.SetFontSize(15)
    pdf.SetXY(55, 199)
    pdf.Cell(nil, "Product title (very good price)")

    pdf.SetFontSize(15)
    pdf.SetXY(451, 317.5)
    pdf.Cell(nil, "2000")

    bytes := pdf.GetBytesPdf()

    ioutil.WriteFile("result.pdf", bytes, 0644)

    pdf.Close()
}


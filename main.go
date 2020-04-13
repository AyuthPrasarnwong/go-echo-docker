package main

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)


type (
	input struct {
		Title   string    `json:"title"`
	}
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "public")

	e.GET("/file", func(c echo.Context) error {
		return c.File("echo.svg")
	})

	e.POST("/genpdf", func(c echo.Context) error {

		input := new(input)

		if err := c.Bind(&input); err != nil {
			return err
		}


		titleStr := input.Title

		pdf := gofpdf.New("P", "mm", "A4", "")

		pdf.SetTitle(titleStr, false)
		pdf.SetAuthor("Jules Verne", false)
		pdf.SetHeaderFunc(func() {
			// Arial bold 15
			pdf.SetFont("Arial", "B", 15)
			// Calculate width of title and position
			wd := pdf.GetStringWidth(titleStr) + 6
			pdf.SetX((210 - wd) / 2)
			// Colors of frame, background and text
			pdf.SetDrawColor(0, 80, 180)
			pdf.SetFillColor(230, 230, 0)
			pdf.SetTextColor(220, 50, 50)
			// Thickness of frame (1 mm)
			pdf.SetLineWidth(1)
			// Title
			pdf.CellFormat(wd, 9, titleStr, "1", 1, "C", true, 0, "")
			// Line break
			pdf.Ln(10)
		})
		pdf.SetFooterFunc(func() {
			// Position at 1.5 cm from bottom
			pdf.SetY(-15)
			// Arial italic 8
			pdf.SetFont("Arial", "I", 8)
			// Text color in gray
			pdf.SetTextColor(128, 128, 128)
			// Page number
			pdf.CellFormat(0, 10, fmt.Sprintf("Page %d", pdf.PageNo()),
				"", 0, "C", false, 0, "")
		})

		pdf.AddPage()

		err := pdf.OutputFileAndClose("test.pdf")

		if err != nil {
			fmt.Println(err)
		}

		return c.File("test.pdf")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
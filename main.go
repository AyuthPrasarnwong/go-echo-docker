package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "public")

	e.GET("/file", func(c echo.Context) error {
		return c.File("echo.svg")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
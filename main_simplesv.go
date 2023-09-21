package main

import (
	"os"

	"github.com/labstack/echo/v4"
)

func main() {

	port := "1323"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	e := echo.New()
	e.Static("/", "")
	e.Logger.Fatal(e.Start(":" + port))
}

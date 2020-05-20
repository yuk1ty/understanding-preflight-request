package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()
}

func handle(c echo.Context) {
	c.Response().Header().Set("some header", "here")
}

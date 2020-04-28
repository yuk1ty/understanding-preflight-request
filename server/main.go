package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http/httputil"
)

func main() {
	fmt.Println("Server starts...")

	e := echo.New()

	e.POST("/test", requestHandle)

	e.Logger.Fatal(e.Start(":8081"))
}

func requestHandle(c echo.Context) error {
	body, err := httputil.DumpRequest(c.Request(), true)
	if err != nil {
		panic(err)
	}

	return c.String(200, string(body))
}

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
	e.OPTIONS("/test", requestHandle2)
	e.Logger.Fatal(e.Start(":8081"))
}
func requestHandle(c echo.Context) error {
	body, err := httputil.DumpRequest(c.Request(), true)
	if err != nil {
		panic(err)
	}
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	return c.String(200, string(body))
}
func requestHandle2(c echo.Context) error {
	body, err := httputil.DumpRequest(c.Request(), true)
	if err != nil {
		panic(err)
	}
	c.Response().Header().Set("Access-Control-Allow-Origin", "*")
	c.Response().Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, OPTIONS")
	c.Response().Header().Set("Access-Control-Allow-Headers", "Content-Type")
	return c.String(200, string(body))
}

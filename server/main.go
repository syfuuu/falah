package main

import (
  "github.com/labstack/echo/v4"
  "net/http"
)

func helloWorld() echo.HandlerFunc {
  return "Hello world!"
} 

func main() {
	e := echo.New()

	e.GET("/", helloWorld)
	e.Start(":8080")
}

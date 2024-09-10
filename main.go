package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)


func main() {
	e := echo.New()
	

	e.GET("/", func(c echo.Context) error {
		fmt.Println(c.Request().Header)
		
		return nil
	})

	e.Logger.Panic(e.Start(fmt.Sprintf(":%s", os.Getenv("PORT"))))
}
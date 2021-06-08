package main

import (
	"github.com/labstack/echo/v4"
	"com/rakib/banking/main/database"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	Router(e)

	database.GetDmManager()
	e.Logger.Fatal(e.Start(":80"))
}

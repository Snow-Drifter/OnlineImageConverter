package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.AddTrailingSlash())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Cross-Origin-Opener-Policy", "same-origin")
			c.Response().Header().Set("Cross-Origin-Embedder-Policy", "credentialless")
			return next(c)
		}
	})

	e.GET("/", func(c echo.Context) error {
		return c.File("index.html")
	})

	e.Static("/", ".")

	port := ":3000"
	e.Logger.Info("Server running on http://localhost" + port)
	e.Logger.Fatal(e.Start(port))
}

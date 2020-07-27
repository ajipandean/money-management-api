package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {
    e := echo.New()

    // Middlewares
    e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
        Format: "method=${method}, uri=${uri}, status=${status}\n",
    }))
    e.Use(middleware.Recover())

    // Routes
    api := e.Group("/api")
    api.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello world")
    })

    e.Logger.Fatal(e.Start("localhost:8080"))
}

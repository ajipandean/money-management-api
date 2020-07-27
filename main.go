package main

import (
    "fmt"
    "net/http"
    "money-management-api/database"
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
    e := echo.New()

    // Database connection
    var err error
    database.DB, err = gorm.Open("mysql", "root:@/db_money_management?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic("MySQL refuse to connect with us :<")
    }

    fmt.Println("MySQL already connected with us :>")

    defer database.DB.Close()

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

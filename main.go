package main

import (
    "fmt"
    "money-management-api/models"
    "money-management-api/database"
    "money-management-api/routes"
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

    database.DB.AutoMigrate(&models.User{}, &models.Wallet{})
    fmt.Println("Database table migrated :>")

    defer database.DB.Close()

    // Middlewares
    e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
        Format: "method=${method}, uri=${uri}, status=${status}\n",
    }))
    e.Use(middleware.Recover())

    // Routes
    routes.Setup(e)

    e.Logger.Fatal(e.Start("localhost:8080"))
}

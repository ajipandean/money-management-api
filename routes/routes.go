package routes

import (
    "github.com/labstack/echo/v4"
    uc "money-management-api/controllers/user"
)

func Setup(e *echo.Echo) {
    v1 := e.Group("/api/v1")

    // User routes
    v1.GET("/users", uc.FetchUsers)
    v1.POST("/users", uc.CreateUser)
    v1.GET("/users/:id", uc.FindUser)
    v1.PUT("/users/:id", uc.UpdateUser)
    v1.DELETE("/users/:id", uc.DeleteUser)
}

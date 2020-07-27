package user_controller

import (
    "net/http"
    "money-management-api/database"
    "github.com/labstack/echo/v4"
)

func FetchUsers(c echo.Context) error {
    return c.String(http.StatusOK, "Users fetched")
}
func FindUser(c echo.Context) error {
    return c.String(http.StatusOK, "Users found")
}
func CreateUser(c echo.Context) error {
    return c.String(http.StatusOK, "Users created")
}
func UpdateUser(c echo.Context) error {
    return c.String(http.StatusOK, "Users updated")
}
func DeleteUser(c echo.Context) error {
    return c.String(http.StatusOK, "Users deleted")
}

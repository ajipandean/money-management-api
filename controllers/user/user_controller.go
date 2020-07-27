package user_controller

import (
    "net/http"
    "github.com/labstack/echo/v4"
    d "money-management-api/database"
    m "money-management-api/models"
)

type M map[string]interface{}

func FetchUsers(c echo.Context) error {
    db := d.DB
    users := new([]m.User)

    db.Preload("Wallets").Find(&users)

    return c.JSON(http.StatusOK, users)
}

func FindUser(c echo.Context) error {
    db := d.DB
    user := new(m.User)

    if res := db.Preload("Wallets").First(user, c.Param("id")); res.RecordNotFound() {
        return c.JSON(http.StatusInternalServerError, M{
            "Message": "User not found",
        })
    }

    return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
    db := d.DB
    user := new(m.User)

    if err := c.Bind(user); err != nil {
        panic(err.Error())
    }

    if res := db.Create(user); res.Error != nil {
        return c.JSON(http.StatusInternalServerError, res.Error)
    }

    return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
    db := d.DB
    user := new(m.User)

    if res := db.First(user, c.Param("id")); res.RecordNotFound() {
        return c.JSON(http.StatusInternalServerError, M{
            "Message": "User not found",
        })
    }

    if err := c.Bind(user); err != nil {
        panic(err.Error())
    }

    db.Save(user);

    return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {
    db := d.DB
    user := new(m.User)

    if res := db.First(user, c.Param("id")); res.RecordNotFound() {
        return c.JSON(http.StatusInternalServerError, M{
            "Message": "User not found",
        })
    }

    db.Delete(user)

    return c.JSON(http.StatusOK, user)
}

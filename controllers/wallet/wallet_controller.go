package wallet_controller

import (
    "net/http"
    "github.com/labstack/echo/v4"
    d "money-management-api/database"
    m "money-management-api/models"
)

type M map[string]interface{}

func FetchWallets(c echo.Context) error {
    db := d.DB
    wallets := new([]m.Wallet)

    db.Find(wallets)

    return c.JSON(http.StatusOK, wallets)
}

func FindWallet(c echo.Context) error {
    db := d.DB
    wallet := new(m.Wallet)

    if res := db.First(wallet, c.Param("id")); res.RecordNotFound() {
        return c.JSON(http.StatusInternalServerError, M{
            "Message": "Wallet not found",
        })
    }

    return c.JSON(http.StatusOK, wallet)
}

func CreateWallet(c echo.Context) error {
    db := d.DB
    wallet := new(m.Wallet)

    if err := c.Bind(wallet); err != nil {
        panic(err.Error())
    }

    if res := db.Create(wallet); res.Error != nil {
        return c.JSON(http.StatusInternalServerError, res.Error)
    }

    return c.JSON(http.StatusOK, wallet)
}

func UpdateWallet(c echo.Context) error {
    db := d.DB
    wallet := new(m.Wallet)

    if res := db.First(wallet, c.Param("id")); res.RecordNotFound() {
        return c.JSON(http.StatusInternalServerError, M{
            "Message": "Wallet not found",
        })
    }

    if err := c.Bind(wallet); err != nil {
        panic(err.Error())
    }

    db.Save(wallet)

    return c.JSON(http.StatusOK, wallet)
}

func DeleteWallet(c echo.Context) error {
    db := d.DB
    wallet := new(m.Wallet)

    if res := db.First(wallet, c.Param("id")); res.RecordNotFound() {
        return c.JSON(http.StatusInternalServerError, M{
            "Message": "Wallet not found",
        })
    }

    db.Delete(wallet)

    return c.JSON(http.StatusOK, wallet)
}

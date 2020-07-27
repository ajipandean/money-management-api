package wallet_controller

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

func FetchWallets(c echo.Context) error {
    return c.String(http.StatusOK, "Wallets fetched")
}

func FindWallet(c echo.Context) error {
    return c.String(http.StatusOK, "Wallets found")
}

func CreateWallet(c echo.Context) error {
    return c.String(http.StatusOK, "Wallets created")
}

func UpdateWallet(c echo.Context) error {
    return c.String(http.StatusOK, "Wallets updated")
}

func DeleteWallet(c echo.Context) error {
    return c.String(http.StatusOK, "Wallets deleted")
}

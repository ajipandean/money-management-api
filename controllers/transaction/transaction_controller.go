package transaction_controller

import (
    "net/http"
    "github.com/labstack/echo/v4"
)

func FetchTransactions(c echo.Context) error {
    return c.String(http.StatusOK, "Transactions fetched")
}

func FindTransaction(c echo.Context) error {
    return c.String(http.StatusOK, "Transaction found")
}

func CreateTransaction(c echo.Context) error {
    return c.String(http.StatusOK, "Transaction created")
}

func UpdateTransaction(c echo.Context) error {
    return c.String(http.StatusOK, "Transaction updated")
}

func DeleteTransaction(c echo.Context) error {
    return c.String(http.StatusOK, "Transaction delete")
}

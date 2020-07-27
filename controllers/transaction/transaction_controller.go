package transaction_controller

import (
    "net/http"
    "github.com/labstack/echo/v4"
    d "money-management-api/database"
    m "money-management-api/models"
)

type M map[string]interface{}

func FetchTransactions(c echo.Context) error {
    db := d.DB
    transactions := new([]m.Transaction)

    db.Find(transactions)

    return c.JSON(http.StatusOK, transactions)
}

func FindTransaction(c echo.Context) error {
    db := d.DB
    transaction := new(m.Transaction)

    if res := db.First(transaction, c.Param("id")); res.RecordNotFound() {
        return c.JSON(http.StatusInternalServerError, M{
            "Message": "Transaction not found",
        })
    }

    return c.JSON(http.StatusOK, transaction)
}

func CreateTransaction(c echo.Context) error {
    db := d.DB
    transaction := new(m.Transaction)

    if err := c.Bind(transaction); err != nil {
        panic(err.Error())
    }

    if res := db.Create(transaction); res.Error != nil {
        return c.JSON(http.StatusInternalServerError, res.Error)
    }

    return c.JSON(http.StatusOK, transaction)
}

func UpdateTransaction(c echo.Context) error {
    db := d.DB
    transaction := new(m.Transaction)

    if res := db.First(transaction, c.Param("id")); res.RecordNotFound() {
        return c.JSON(http.StatusInternalServerError, M{
            "Message": "Transaction not found",
        })
    }

    if err := c.Bind(transaction); err != nil {
        panic(err.Error())
    }

    db.Save(transaction)

    return c.JSON(http.StatusOK, transaction)
}

func DeleteTransaction(c echo.Context) error {
    db := d.DB
    transaction := new(m.Transaction)

    if res := db.First(transaction, c.Param("id")); res.RecordNotFound() {
        return c.JSON(http.StatusInternalServerError, M{
            "Message": "Transaction not found",
        })
    }

    db.Delete(transaction)

    return c.JSON(http.StatusOK, transaction)
}

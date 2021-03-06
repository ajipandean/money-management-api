package routes

import (
    "github.com/labstack/echo/v4"
    uc "money-management-api/controllers/user"
    wc "money-management-api/controllers/wallet"
    tc "money-management-api/controllers/transaction"
)

func Setup(e *echo.Echo) {
    v1 := e.Group("/api/v1")

    // User routes
    v1.GET("/users", uc.FetchUsers)
    v1.POST("/users", uc.CreateUser)
    v1.GET("/users/:id", uc.FindUser)
    v1.PUT("/users/:id", uc.UpdateUser)
    v1.DELETE("/users/:id", uc.DeleteUser)

    // Wallet routes
    v1.GET("/wallets", wc.FetchWallets)
    v1.POST("/wallets", wc.CreateWallet)
    v1.GET("/wallets/:id", wc.FindWallet)
    v1.PUT("/wallets/:id", wc.UpdateWallet)
    v1.DELETE("/wallets/:id", wc.DeleteWallet)

    // Transaction routes
    v1.GET("/transactions", tc.FetchTransactions)
    v1.POST("/transactions", tc.CreateTransaction)
    v1.GET("/transactions/:id", tc.FindTransaction)
    v1.PUT("/transactions/:id", tc.UpdateTransaction)
    v1.DELETE("/transactions/:id", tc.DeleteTransaction)
}

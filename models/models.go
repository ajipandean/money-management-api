package models

import (
    "time"
    "golang.org/x/crypto/bcrypt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

// User model
// Has One2Many relation to Wallet
type User struct {
    gorm.Model
    Email    string `gorm:"UNIQUE"`
    Username string `gorm:"not null"`
    Password string `gorm:"not null"`
    Wallets  []Wallet
}
// Hash User password before
// it being saved to database
func (u *User) BeforeSave() error {
    hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    u.Password = string(hashed)

    return nil
}

// Wallet model
// Belongs to User
// Has One2Many relation to Transaction
type Wallet struct {
    gorm.Model
    Name         string `gorm:"not null"`
    Currency     string `gorm:"size:10"`
    Balance      int    `gorm:"not null"`
    UserID       uint   `gorm:"not null"`
    Transactions []Transaction
}

// Transaction model
// Belongs to Wallet
type Transaction struct {
    gorm.Model
    Date        time.Time `gorm:"not null"`
    Description string    `gorm:"type:text;"`
    Amount      int       `gorm:"not null"`
    WalletID    uint      `gorm:"not null"`
}

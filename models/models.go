package models

import (
    "golang.org/x/crypto/bcrypt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

// User model
// Has One2Many relation to Wallet
type User struct {
    gorm.Model
    Email    string `gorm:"UNIQUE"`
    Username string
    Password string
    Wallets  []Wallet
}
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
type Wallet struct {
    gorm.Model
    Name     string
    Currency string `gorm:"size:10"`
    Balance  int
    UserID   uint
    User     User
}

package models

import (
    "golang.org/x/crypto/bcrypt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
    gorm.Model
    Email    string `gorm:"UNIQUE"`
    Username string
    Password string
}
func (u *User) BeforeSave() error {
    hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    u.Password = string(hashed)

    return nil
}

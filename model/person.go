package model

import (
	"errors"
	"fmt"
	"matmuh/helper/message"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	PersonID      int
	Name          string
	LastName      string
	Email         string
	Password      string
	PasswordAgain string
	DateOfBirth   string
}

func (p Person) RegisterPerson(db *gorm.DB) error {
	if p.Password != p.PasswordAgain {
		fmt.Println("Sifreler uyumlu degil")
		return errors.New(message.PasswordErrorMessage)
	}
	result := db.Where(&Person{Email: p.Email}).Find(&Person{})
	if result.RowsAffected > 0 {
		fmt.Println("Bu mail adresi daha once kullanilmis")
		return errors.New(message.EmailErrorMessage)
	}
	db.Create(&p)
	return nil
}

package main

import (
	"matmuh/helper/db"
	"matmuh/model"
)

func main() {
	DB := db.InitDB()
	DB.AutoMigrate(&model.Person{})

	var p model.Person = model.Person{
		PersonID:      1,
		Name:          "Osman Bahadir",
		LastName:      "Avci",
		Email:         "deneme@mail.com",
		Password:      "123123",
		PasswordAgain: "123123",
		DateOfBirth:   "05/10/2001",
	}
	p.RegisterPerson(DB)
}

/*
model
v
c

email daha once kullanulmis mi
password uyuyor mu

*/

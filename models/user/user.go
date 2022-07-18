package models

import (
	"learn-gin/utils"
	"time"
)

type User struct {
	ID        uint32    `gorm:"column:id;primaryKey" json:"id,omitempty"`
	FirstName string    `gorm:"column:first_name;type:varchar(50)" json:"first_name"`
	LastName  string    `gorm:"column:last_name;type:varchar(50)" json:"last_name"`
	Gender    string    `gorm:"column:gender;type:char(10)" json:"gender"`
	BirthDate string    `gorm:"column:birth_date;type:date" json:"birth_date"`
	Email     string    `gorm:"column:email;unique;type:varchar(100)" json:"email"`
	Password  string	`gorm:"column:password;type:varchar(255)" json:"password"`
	Phone     string    `gorm:"column:phone;type:char(20)" json:"phone"`
	IsActive  *bool     `gorm:"column:is_active;default:true" json:"is_active,omitempty"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}

func (user *User) Validation() (strMessage string) {
	if !utils.IsAlphabet(user.FirstName) {
		strMessage = "First Name must only contains alphabet and can't be empty!"
		return
	}
	if !utils.IsAlphabet(user.LastName) {
		strMessage = "Last Name must only contains alphabet and can't be empty!"
		return
	}

	if !utils.IsBirthDateValid(user.BirthDate) {
		strMessage = "Birth Date format data is not valid! Accepted format: YYYY-MM-DD"
		return
	}
	if !utils.IsEmailValid(user.Email) {
		strMessage = "Email address is not valid!"
		return
	}

	if !utils.IsPhoneValid(user.Phone) {
		strMessage = "Phone Number is not valid. Phone Number only contains numeric. min 7 digit, max 16 digit!"
		return
	}
	return
}

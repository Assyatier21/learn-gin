package admin

import (
	"learn-gin/utils"
	"time"
)

type Admin struct {
	ID           uint32    `gorm:"column:id;primaryKey" json:"id,omitempty"`
	FirstName    string    `gorm:"column:first_name;type:varchar(50)" json:"first_name"`
	LastName     string    `gorm:"column:last_name;type:varchar(50)" json:"last_name"`
	Email        string    `gorm:"column:email;unique;type:varchar(100)" json:"email"`
	Password     string    `gorm:"column:password;type:varchar(255)" json:"password"`
	Phone        string    `gorm:"column:phone;type:char(20)" json:"phone"`
	IsActive     *bool     `gorm:"column:is_active;default:true" json:"is_active,omitempty"`
	ID_Previlege uint32    `gorm:"column:id_previlege;type:int(10)" json:"id_previlege"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at,omitempty"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at,omitempty"`
}

func (adm *Admin) Validation(strMessage string) {
	if !utils.IsAlphabet(adm.FirstName) {
		strMessage = "First Name must only contains alphabet and can't be empty!"
		return
	}
	if !utils.IsAlphabet(adm.LastName) {
		strMessage = "Last Name must only contains alphabet and can't be empty!"
		return
	}
	if !utils.IsEmailValid(adm.Email) {
		strMessage = "Email address is not valid!"
		return
	}
	if !utils.IsPhoneValid(adm.Phone) {
		strMessage = "Phone Number is not valid. Phone Number only contains numeric. min 7 digit, max 16 digit!"
		return
	}
	return
}

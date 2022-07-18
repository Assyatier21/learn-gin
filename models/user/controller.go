package models

import (
	"learn-gin/datasources/mysql"
	"log"
)

func CreateUser(user *User) (err error) {
	if err = mysql.Connection.Create(user).Error; err != nil {
		return err
	}
	return nil
}
func UpdateUser(user *User) (err error) {
	log.Println(user)
	mysql.Connection.Updates(user)
	return nil
}

func GetUserByID(user *User, id uint32) (err error) {
	if err = mysql.Connection.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func GetAllUsers(user *[]User) (err error) {
	if err = mysql.Connection.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *User, id uint32) (err error) {
	mysql.Connection.Where("id = ?", id).Delete(user)
	return nil
}

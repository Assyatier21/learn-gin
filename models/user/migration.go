package models

import (
	"learn-gin/datasources/mysql"
	"log"
)

func Migration() {
	if !mysql.Connection.Migrator().HasTable(&User{}) {
		err := mysql.Connection.Migrator().CreateTable(&User{})
		if err != nil {
			log.Println(err.Error())
		}
	}
}

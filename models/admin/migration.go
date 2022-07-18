package admin

import (
	"learn-gin/datasources/mysql"
	"log"
)

func Migration() {
	if !mysql.Connection.Migrator().HasTable(&Admin{}) {
		err := mysql.Connection.Migrator().CreateTable(&Admin{})
		if err != nil {
			log.Println(err.Error())
		}
	}
}

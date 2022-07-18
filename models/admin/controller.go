package admin

import (
	"learn-gin/datasources/mysql"
)

func CreateAdmin(adm *Admin) (err error) {
	if err = mysql.Connection.Create(adm).Error; err != nil {
		return nil
	}
	return err
}
func DeleteAdmin(adm *Admin, id uint32) (err error) {
	mysql.Connection.Where("id = ?", id).Delete(adm)
	return nil
}

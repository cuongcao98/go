package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Conn is connect database
func Conn() *gorm.DB {
	// Openning file
	db, _ := gorm.Open("mysql", "root:root@/user?charset=utf8&parseTime=True&loc=Local")

	return db
}

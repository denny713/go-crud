package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func SetupKoneksi() *Database {
	dbase := Database{
		Host: "localhost",
		Port: 3306,
		User: "root",
		Pass: "p@ssw0rd",
		Name: "crud_demo",
	}
	return &dbase
}

func KoneksiLink(dbase *Database) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", dbase.User, dbase.Pass, dbase.Host, dbase.Port, dbase.Name)
}

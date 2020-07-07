package main

import (
	"Github/go-crud/config"
	"Github/go-crud/model"
	"Github/go-crud/route"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var err error

func main() {
	SetupDb()
	x := route.SetupRouter()
	x.Run(":7081")
}

func SetupDb() {
	config.DB, err = gorm.Open("mysql", config.KoneksiLink(config.SetupKoneksi()))
	if err != nil {
		fmt.Println("Error : ", err)
	}
	config.DB.AutoMigrate(&model.User{})
}

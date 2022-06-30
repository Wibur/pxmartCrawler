package databases

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	UserName     string = "root"
	Password     string = "123456"
	Addr         string = "127.0.0.1"
	Port         int    = 3306
	Database     string = "pxmart"
	MaxLifetime  int    = 10
	MaxOpenConns int    = 10
	MaxIdleConns int    = 10
)

var DBConnect *gorm.DB
var err error

func MySqlConnect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", UserName, Password, Addr, Port, Database)
	log.Println(dsn)
	DBConnect, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("MySQL Connect success")
}

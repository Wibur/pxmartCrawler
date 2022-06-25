package databases

import (
	"database/sql"
	"fmt"
	"time"

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

var db *sql.DB

func MySqlConnect() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", UserName, Password, Addr, Port, Database)
	connector, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("connection mysql failed:", err)
		return
	}
	db, err = connector.DB()
	defer db.Close()
	if err != nil {
		fmt.Println("get DB failed:", err)
		return
	}
	// 連線可複用最大時間
	db.SetConnMaxLifetime(time.Duration(MaxLifetime) * time.Second)
	// 空閒池中最大數量
	db.SetMaxIdleConns(MaxIdleConns)
	// 數據庫連接最大數量
	db.SetMaxOpenConns(MaxOpenConns)
}

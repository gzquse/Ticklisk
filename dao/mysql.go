package dao

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	dsn := "root@(127.0.0.1:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return DB.DB().Ping()
}

func Close() {
	err := DB.Close()
	if err != nil {
		return
	}
}

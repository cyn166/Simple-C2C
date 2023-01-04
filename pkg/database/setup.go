package database

import (
	"github.com/cyn166/Simple-C2C/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := "root:@tcp(127.0.0.1:3306)/c2c?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Migrate(db)
	DB = db
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB {
	return DB
}

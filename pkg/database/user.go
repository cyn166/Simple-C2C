package database

import (
	"fmt"

	"github.com/cyn166/Simple-C2C/pkg/models"
)

func GetUser(username string, password string) (models.User, error) {
	var user models.User
	db := GetDB()
	err := db.Where("username = ?", username).Where("password = ?", password).First(&user).Error
	return user, err
}

func CheckUsernameExists(username string) bool {
	var user models.User
	db := GetDB()
	err := db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return false
	}
	fmt.Println(user.Username)
	return true
}

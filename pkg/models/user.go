package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Vip      bool
	Plan     string
	Password string
	Admin    bool
	Mod      bool
	Support  bool
	Resseler bool
}

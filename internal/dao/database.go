package dao

import (
	initialization "github.com/isHuangxin/tiktok-backend/init"
	"gorm.io/gorm"
)

var db *gorm.DB

func DataBaseInitialization() {
	db = initialization.GetDB()
}

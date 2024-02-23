package config

import (
	domain "dating-app/service-user/domain"
	sharedConfig "dating-app/shared/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitDB() *gorm.DB {

	db := sharedConfig.InitDB()

	createTable(db)
	migrateDDL(db)

	return db
}

func createTable(db *gorm.DB) {
	db.CreateTable(&domain.User{})
}

func migrateDDL(db *gorm.DB) {
	db.AutoMigrate(&domain.User{})
}

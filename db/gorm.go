package db

import (
	"log"

	"wire-gorm-server/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectGORM() *gorm.DB {
	dsn := "root:password@tcp(localhost:3306)/testdb?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	if err := db.AutoMigrate(&model.User{}); err != nil {
		log.Fatal("failed to migrate database:", err)
	}

	return db
}

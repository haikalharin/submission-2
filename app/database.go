package app

import (
	"fmt"
	"submission-2/helper"
	"submission-2/model/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host   = "localhost"
	user   = "postgres"
	pass   = "root"
	port   = "5433"
	dbName = "submission_2"
)

func NewDB() *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, dbName, port)
	db, err := gorm.Open(postgres.Open(config))
	helper.PanicIfError(err)

	err = db.AutoMigrate(&domain.Order{}, &domain.Item{})
	helper.PanicIfError(err)

	return db
}

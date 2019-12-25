package main

import (
	"startupbruhs.github.io/bookstore/db/model/entity"
	"startupbruhs.github.io/bookstore/server"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	db, DBErr := gorm.Open(
		os.Getenv("DB_DRIVE"),
		os.Getenv("DB_USERNAME")+":"+os.Getenv("DB_PASSWORD")+"@tcp("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local",
	)
	if DBErr != nil {
		panic("DBErr.Error()" + DBErr.Error())
	}

	db.LogMode(true)

	db.Debug().AutoMigrate(&entity.User{}, &entity.Author{}, &entity.Book{})
	server.Start(db)

	defer db.Close()
}

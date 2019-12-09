package repository

import (
	"bookstore/db/model/entity"

	"github.com/jinzhu/gorm"
)

type UserRepo struct {
	Conn *gorm.DB
}

func (repo *UserRepo) All() []entity.User {
	var users []entity.User
	repo.Conn.Find(&users)
	return users
}

func (repo *UserRepo) TestCreate() {
	repo.Conn.Create(&entity.User{
		Name:     "bernard",
		Email:    "bernardkllogjri@gmail.com",
		Avatar:   "https://www.google.com/url?sa=i&source=images&cd=&ved=2ahUKEwj1nse2vafmAhVFZVAKHSG4BT0QjRx6BAgBEAQ&url=https%3A%2F%2Fcollider.com%2Favatar-sequels-release-dates-delayed-for-star-wars-movies%2F&psig=AOvVaw2atMUo9m1513hVpDgS3uKK&ust=1575943562754259",
		Password: "password",
	})
}

func InitUserRepo(database *gorm.DB) *UserRepo {
	return &UserRepo{Conn: database}
}

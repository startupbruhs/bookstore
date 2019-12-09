package server

import (
	"bookstore/routes"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Start(Connection *gorm.DB) {
	r := gin.Default()

	routes.User(r, Connection)
	routes.Book(r, Connection)

	r.Run()
}

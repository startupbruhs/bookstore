package routes

import (
	"bookstore/repository"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// User Routes
func User(r *gin.Engine, db *gorm.DB) {
	user := repository.InitUserRepo(db)
	baseURL := "/users"

	r.GET(baseURL+"/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"data":    user.All(),
		})
	})

	r.GET(baseURL+"/create", func(c *gin.Context) {
		user.TestCreate()
		c.JSON(200, gin.H{
			"status": 200,
		})
	})

}

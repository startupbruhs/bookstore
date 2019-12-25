package server

import (
	"startupbruhs.github.io/bookstore/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//Start Function to start the server
func Start(Connection *gorm.DB) {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "http://localhost:3000"
		// },
		MaxAge: 12 * time.Hour,
	}))
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": 200,
		})
	})

	routes.User(router, Connection)
	routes.Book(router, Connection)

	router.Run(":9000")
}

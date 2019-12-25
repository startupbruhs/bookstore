package routes

import (
	"startupbruhs.github.io/bookstore/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Book Routes
func Book(r *gin.Engine, db *gorm.DB) {
	user := repository.InitUserRepo(db)
	baseURL := "/books"

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

	r.GET(baseURL+"/create/from-json", func(c *gin.Context) {
		jsonFile, err := os.Open("db.json")
		if err != nil {
			fmt.Println(err)
		}
		byteValue, _ := ioutil.ReadAll(jsonFile)
		var result []interface{}

		json.Unmarshal([]byte(byteValue), &result)

		c.JSON(200, gin.H{
			"status": 200,
			"books":  result,
		})
	})
}

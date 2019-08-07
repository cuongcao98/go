package controller

import (
	"project2/database"

	"github.com/gin-gonic/gin"
)

//ReadAll is select db
func ReadAll(c *gin.Context) {
	// Connection to the database
	db := database.Conn()
	// Close connection database
	defer db.Close()

	var info []Infos
	// SELECT * FROM users
	db.Find(&info)

	// Display JSON result
	c.JSON(200, info)

}

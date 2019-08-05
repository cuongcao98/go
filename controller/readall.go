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

	var thongtin []Thongtins
	// SELECT * FROM users
	db.Find(&thongtin)

	// Display JSON result
	c.JSON(200, thongtin)

}

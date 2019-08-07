package controller

import (
	"project2/database"

	"github.com/gin-gonic/gin"
)

//Read is select
func Read(c *gin.Context) {
	// Connection to the database
	db := database.Conn()
	// Close connection database
	defer db.Close()

	id := c.Params.ByName("id")
	var info Infos
	// SELECT * FROM users WHERE id = 1;
	db.First(&info, id)

	// Display JSON result
	c.JSON(200, info)

}

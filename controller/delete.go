package controller

import (
	"project2/database"

	"github.com/gin-gonic/gin"
)

// Delete func ..
func Delete(c *gin.Context) {
	db := database.Conn()

	id := c.Params.ByName("id")
	var info Infos

	db.Delete(&info, "id = ?", id)
	c.JSON(200, gin.H{
		"messages": "deleted",
	})

	defer db.Close()
}

package controller

import (
	"project/database"

	"github.com/gin-gonic/gin"
)

// Delete func ..
func Delete(c *gin.Context) {
	db := database.Conn()

	delete, err := db.Prepare("DELETE FROM thongtins WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	delete.Exec(c.Param("id"))
	c.JSON(200, gin.H{
		"messages": "deleted",
	})

	defer db.Close()
}

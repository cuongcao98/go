package controller

import (
	"project/database"
	"time"

	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	db := database.Conn()
	type UpdatePost struct {
		Username    string    `form:"username" json:"username" binding:"required"`
		Password    string    `form:"password" json:"password" binding:"required"`
		Email       string    `form:"email" json:"email" binding:"required"`
		Birth       time.Time `form:"birth" json:"birth" binding:"required"`
		Sex         string    `form:"sex" json:"sex" binding:"required"`
		Phone       string    `form:"phone" json:"phone" binding:"required"`
		National_id int       `form:"national_id" json:"national_id" binding:"required"`
		Height      float64   `form:"height" json:"height" binding:"required"`
	}

	var json UpdatePost
	if err := c.ShouldBindJSON(&json); err == nil {
		edit, err := db.Prepare("UPDATE thongtins SET username=?, password=?, email=?, birth=?, sex=?, phone=?, national_id=?, height=? WHERE id = " + c.Param("id"))
		if err != nil {
			panic(err.Error())
		}
		edit.Exec(json.Username, json.Password, json.Email, json.Birth, json.Sex, json.Phone, json.National_id, json.Height)

		c.JSON(200, gin.H{
			"messages": "edited",
		})
	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}
	defer db.Close()
}

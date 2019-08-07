package controller

import (
	"project2/database"
	"time"

	"github.com/gin-gonic/gin"
)

//Update is ..
func Update(c *gin.Context) {
	db := database.Conn()

	type UpdatePost struct {
		Username   string
		Password   string
		Email      string
		Birth      time.Time
		Sex        string
		Phone      string
		NationalID int
		Height     float64
	}
	var info UpdatePost

	if err := c.ShouldBindJSON(&info); err == nil {
		info = UpdatePost{
			Username:   info.Username,
			Password:   GetMD5Hash(info.Password),
			Email:      info.Email,
			Birth:      info.Birth,
			Sex:        info.Sex,
			Phone:      info.Phone,
			NationalID: info.NationalID,
			Height:     info.Height,
		}
		db.Table("infos").Where("ID = ?", c.Param("id")).Updates(&info)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": err,
			})
		}
		c.JSON(200, gin.H{
			"messages": "thanh cong",
		})

	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	defer db.Close()

}

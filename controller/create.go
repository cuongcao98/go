package controller

import (
	"crypto/md5"
	"encoding/hex"
	"project2/database"
	"time"

	"github.com/gin-gonic/gin"
)

//Create new user
func Create(c *gin.Context) {
	db := database.Conn()
	defer db.Close()
	type Info struct {
		Username   string
		Password   string
		Email      string
		Birth      time.Time
		Sex        string
		Phone      string
		NationalID int
		Height     float64
	}

	var info Info

	if err := c.ShouldBindJSON(&info); err == nil {
		info = Info{Username: info.Username,
			Password:   GetMD5Hash(info.Password),
			Email:      info.Email,
			Birth:      info.Birth,
			Sex:        info.Sex,
			Phone:      info.Phone,
			NationalID: info.NationalID,
			Height:     info.Height}
		db.Create(&info)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": err,
			})
		}
		c.JSON(200, gin.H{
			"messages": info,
		})

	} else {
		c.JSON(500, gin.H{"error": err.Error()})
	}

	defer db.Close()

}

//GetMD5Hash is ..
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

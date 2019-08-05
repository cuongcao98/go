package controller

import (
	"crypto/md5"
	"encoding/hex"
	"project2/database"
	"time"

	"github.com/gin-gonic/gin"
)

//Create is ..
func Create(c *gin.Context) {
	db := database.Conn()
	defer db.Close()
	type Thongtin struct {
		Username   string    `form:"username" json:"username" binding:"required"`
		Password   string    `form:"password" json:"password" binding:"required"`
		Email      string    `form:"email" json:"email" binding:"required"`
		Birth      time.Time `form:"birth" json:"birth" binding:"required"`
		Sex        string    `form:"sex" json:"sex" binding:"required"`
		Phone      string    `form:"phone" json:"phone" binding:"required"`
		NationalID int       `form:"national_id" json:"nationalID" binding:"required"`
		Height     float64   `form:"height" json:"height" binding:"required"`
	}

	var thongtin Thongtin
	// c.JSON(201, gin.H{"success": thongtin})

	if err := c.ShouldBindJSON(&thongtin); err == nil {
		thongtin = Thongtin{Username: thongtin.Username, Password: GetMD5Hash(thongtin.Password), Email: thongtin.Email, Birth: thongtin.Birth, Sex: thongtin.Sex, Phone: thongtin.Phone, NationalID: thongtin.NationalID, Height: thongtin.Height}
		db.Create(&thongtin)
		if err != nil {
			c.JSON(500, gin.H{
				"messages": err,
			})
		}
		c.JSON(200, gin.H{
			"messages": thongtin,
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

package main

import (
	"project2/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(Cors())
	v := r.Group("api")
	{
		v.GET("/alluser", controller.ReadAll)
		v.GET("/user/:id", controller.Read)
		v.POST("/create", controller.Create)
		v.PUT("/update/:id", controller.Update)
		v.DELETE("/delete/:id", controller.Delete)
	}
	r.Run(":8080")
}

//Cors is ..
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

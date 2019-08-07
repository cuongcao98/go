package main

import (
	"project2/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v := r.Group("api")
	{
		r.Use(Cors())
		v.GET("/alluser", controller.ReadAll)
		v.GET("/user/:id", controller.Read)
		v.POST("/create", controller.Create)
		v.POST("/update/:id", controller.Update)
		v.DELETE("/delete/:id", controller.Delete)
	}
	r.Run(":8080")
}

//Cors is ..
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "http://localhost:4200/")
		c.Next()
	}
}

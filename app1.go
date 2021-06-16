package main

import "github.com/gin-gonic/gin"
import "time"
import "fmt"
// import "context"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		time.Sleep(1 * time.Second)
		fmt.Println("Done after 1 sec")
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8090")
}
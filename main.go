package main

import "github.com/gin-gonic/gin"
import "time"
import "fmt"
import "net/http"
import "context"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		go func(){
			time.Sleep(2 * time.Second)
			fmt.Println("Done")
		}()
		req, err := http.NewRequest("GET","http://localhost:8090/ping",nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		// nc := context.WithValue(c.Request.Context(),"t", 1)
		nc := context.WithValue(context.Background(),"t", 1)
		// req = req.WithContext(c.Request.Context())
		req = req.WithContext(nc)
		client := http.DefaultClient
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%v\n", res.StatusCode)
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
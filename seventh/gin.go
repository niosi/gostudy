package main

import "github.com/gin-gonic/gin"
import (
	_ "github.com/golang/protobuf/proto"
	"net/http"
	"fmt"
	"time"
	"log"
)

type IdCard struct {
	Id     int
	CardId string
}

func midd() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware")
		c.Set("request", "clinet_request")
		c.Next()
		fmt.Println("after middleware")
	}
}

func main() {
	r := gin.Default()
	r.Use(midd())
	{
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		r.GET("/user/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.XML(http.StatusOK, gin.H{"username": name, "IdCard": IdCard{1, "9999"}, "age": 127})
		})
		r.GET("/url/:urladdress", func(c *gin.Context) {
			url := c.Param("urladdress")
			fmt.Println(url)
			c.Redirect(http.StatusMovedPermanently, "http://"+url)
		})

		r.GET("/sync", func(c *gin.Context) {
			time.Sleep(5 * time.Second)
			log.Println("Done! in path" + c.Request.URL.Path)
		})

		r.GET("/async", func(c *gin.Context) {
			cCp := c.Copy()
			go func() {
				time.Sleep(5 * time.Second)
				log.Println("Done! in path" + cCp.Request.URL.Path)
			}()
			log.Println(c.GetQueryArray("type"))
		})
	}

	r.Run(":8090") //default  listen and serve on 0.0.0.0:8080
}

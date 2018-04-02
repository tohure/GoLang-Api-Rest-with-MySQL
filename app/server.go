package main

import "fmt"
import "github.com/gin-gonic/gin"
import "net/http"

func main() {

    fmt.Println("Running Server - Tohure")

    r := gin.Default()

    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello Tohure")
    })

    r.GET("/ping", func(c *gin.Context) {
        c.String(http.StatusOK, "pong")
    })

    r.Run(":8080")
}
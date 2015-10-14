package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("*.tmpl")

    router.GET("/user/:name/greet.html", func(c *gin.Context) {
        c.HTML(http.StatusOK, "ginExample.tmpl", gin.H{
            "name": c.Param("name")},
        )
    })
    router.GET("/user/:name/greet.json", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "greeting": fmt.Sprintf("Hello %s.", c.Param("name")),
        })
    })

    router.Run(":8080")
}

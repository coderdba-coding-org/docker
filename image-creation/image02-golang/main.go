package main

import (
        "github.com/gin-gonic/gin"
        cors "github.com/rs/cors/wrapper/gin"
        "log"
        //"net/http"
        //"encoding/json"
        //"fmt"
        //"os"
        //"io/ioutil"
        //"strings"
        //"errors"
        //client "github.com/influxdata/influxdb1-client/v2"

        //"time"
        //"errors"
        //"strconv"
        //"github.com/gin-contrib/cors"
        //client "github.com/influxdata/influxdb/client/v2"
)


func main() {

        log.Println("main(): Entering main()")
        log.Println("main(): Starting web server")

        StartWebServer()
}

func StartWebServer() {

        // ROUTER WITH CUSTOM SETTINGS 2 (with "github.com/rs/cors/wrapper/gin")
        router := gin.Default()
        router.Use(cors.AllowAll())

        // homepage
        router.GET("/", func(c *gin.Context) {
                c.JSON(200, gin.H{
                     "message": "Welcome!",
                })
        })

        // start the web server
        router.Run(":8081")

}

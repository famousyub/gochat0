package main


import (
    "github.com/gin-gonic/gin"
    "github.com/famousyub/gochat0/ws"
)
//https://github.com/famousyub/gochat0
//server
func main() {

    //gin.SetMode(gin.ReleaseMode）// online environment

    go ws.Manager.Start()
    r := gin.Default()
    r.GET("/ws",ws.WsHandler)
    r.GET("/pong", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run(":8282") // listen and serve on 0.0.0.0:8080
}

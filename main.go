package main

import (
    "github.com/gin-gonic/gin"
    "github.com/marcellinoknl/PPL-FinalProject/models"
)
func main() {
    models.ConnectDB()

    r := gin.Default()
	
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    r.Run(":8080")
}

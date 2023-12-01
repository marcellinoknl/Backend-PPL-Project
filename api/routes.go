package api

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
   r.GET("/api/hello", func(c *gin.Context) {
       c.JSON(200, gin.H{"message": "Hello, Golang!"})
   })
}

package main

import (
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "healthy", "service": "09-chat-service"})
    })

    log.Printf("Starting 09-chat-service on port 5009")
    r.Run(":5009")
}

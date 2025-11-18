package main

import (
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "healthy", "service": "12-sms-service"})
    })

    log.Printf("Starting 12-sms-service on port 5012")
    r.Run(":5012")
}

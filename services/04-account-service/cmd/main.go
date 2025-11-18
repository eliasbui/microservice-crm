package main

import (
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "healthy", "service": "04-account-service"})
    })

    log.Printf("Starting 04-account-service on port 5004")
    r.Run(":5004")
}

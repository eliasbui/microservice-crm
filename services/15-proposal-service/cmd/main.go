package main

import (
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "healthy", "service": "15-proposal-service"})
    })

    log.Printf("Starting 15-proposal-service on port 5015")
    r.Run(":5015")
}

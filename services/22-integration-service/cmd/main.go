package main

import (
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "healthy", "service": "22-integration-service"})
    })

    log.Printf("Starting 22-integration-service on port 5022")
    r.Run(":5022")
}

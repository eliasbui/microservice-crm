package main

import (
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "healthy", "service": "18-lead-service"})
    })

    log.Printf("Starting 18-lead-service on port 5018")
    r.Run(":5018")
}

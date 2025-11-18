package main

import (
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "healthy", "service": "05-opportunity-service"})
    })

    log.Printf("Starting 05-opportunity-service on port 5005")
    r.Run(":5005")
}

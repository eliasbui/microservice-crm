package main

import (
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "healthy", "service": "20-reporting-service"})
    })

    log.Printf("Starting 20-reporting-service on port 5020")
    r.Run(":5020")
}

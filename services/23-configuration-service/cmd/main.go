package main

import (
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    r := gin.Default()

    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "healthy", "service": "23-configuration-service"})
    })

    log.Printf("Starting 23-configuration-service on port 5023")
    r.Run(":5023")
}

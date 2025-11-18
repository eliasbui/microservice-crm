package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
}

func main() {
	port := os.Getenv("INTEGRATION_SERVICE_PORT")
	if port == "" {
		port = "8086"
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(ginLogger())

	// Health endpoints
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy", "service": "integration-service"})
	})
	router.GET("/ready", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ready", "service": "integration-service"})
	})
	router.GET("/live", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "alive", "service": "integration-service"})
	})
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// API routes
	api := router.Group("/api/integrations")
	{
		api.POST("/erp/sync", syncERP)
		api.POST("/email-marketing/sync", syncEmailMarketing)
		api.POST("/accounting/sync", syncAccounting)
		api.GET("/status/:integrationId", getIntegrationStatus)
	}

	log.Infof("Integration Service starting on port %s", port)
	router.Run(fmt.Sprintf(":%s", port))
}

func ginLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		log.WithFields(logrus.Fields{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
			"status": c.Writer.Status(),
		}).Info("Request processed")
	}
}

func syncERP(c *gin.Context) {
	log.Info("Syncing with ERP system")
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "ERP sync initiated",
		"syncId":  "erp_sync_123",
	})
}

func syncEmailMarketing(c *gin.Context) {
	log.Info("Syncing with email marketing platform")
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Email marketing sync initiated",
		"syncId":  "email_sync_123",
	})
}

func syncAccounting(c *gin.Context) {
	log.Info("Syncing with accounting system")
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Accounting sync initiated",
		"syncId":  "accounting_sync_123",
	})
}

func getIntegrationStatus(c *gin.Context) {
	integrationId := c.Param("integrationId")
	log.Infof("Getting status for integration: %s", integrationId)
	c.JSON(http.StatusOK, gin.H{
		"integrationId": integrationId,
		"status":        "running",
		"progress":      75,
	})
}

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/crm/chat-ai-service/config"
	"github.com/crm/chat-ai-service/handler"
	"github.com/crm/chat-ai-service/service"
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
	log.Info("Starting Chat & AI Service...")

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize services
	aiService, err := service.NewAIService(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize AI service: %v", err)
	}

	chatService, err := service.NewChatService(cfg, aiService)
	if err != nil {
		log.Fatalf("Failed to initialize chat service: %v", err)
	}

	// Initialize message queue
	messageQueue, err := service.NewMessageQueue(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize message queue: %v", err)
	}
	defer messageQueue.Close()

	// Start consuming messages
	go messageQueue.StartConsuming()

	// Setup router
	router := setupRouter(cfg, chatService, aiService)

	// Create HTTP server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ServerPort),
		Handler: router,
	}

	// Start server in a goroutine
	go func() {
		log.Infof("Server starting on port %s", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info("Shutting down server...")

	// Graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Info("Server exited")
}

func setupRouter(cfg *config.Config, chatService *service.ChatService, aiService *service.AIService) *gin.Engine {
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(ginLogger())

	// CORS middleware
	router.Use(corsMiddleware())

	// Health check endpoints
	router.GET("/health", healthCheck)
	router.GET("/ready", readinessCheck)
	router.GET("/live", livenessCheck)

	// Prometheus metrics
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// API routes
	api := router.Group("/api")
	{
		// Chat endpoints
		chatHandler := handler.NewChatHandler(chatService)
		api.GET("/ws", chatHandler.HandleWebSocket)
		api.POST("/messages", chatHandler.SendMessage)
		api.GET("/messages/:sessionId", chatHandler.GetMessageHistory)

		// AI endpoints
		aiHandler := handler.NewAIHandler(aiService)
		api.POST("/ai/chat", aiHandler.Chat)
		api.POST("/ai/analyze-sentiment", aiHandler.AnalyzeSentiment)
		api.POST("/ai/classify-intent", aiHandler.ClassifyIntent)
		api.POST("/ai/generate-email", aiHandler.GenerateEmail)
	}

	return router
}

func ginLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		duration := time.Since(startTime)

		log.WithFields(logrus.Fields{
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"status":     c.Writer.Status(),
			"duration":   duration.Milliseconds(),
			"ip":         c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}).Info("Request processed")
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"service": "chat-ai-service",
	})
}

func readinessCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ready",
		"service": "chat-ai-service",
	})
}

func livenessCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "alive",
		"service": "chat-ai-service",
	})
}

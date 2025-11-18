package main

import (
    "context"
    "fmt"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/gocql/gocql"
)

type Activity struct {
    CustomerID   int64     `json:"customer_id"`
    ActivityTime time.Time `json:"activity_time"`
    ActivityID   string    `json:"activity_id"`
    ActivityType string    `json:"activity_type"`
    Description  string    `json:"description"`
    UserID       string    `json:"user_id"`
    Metadata     string    `json:"metadata"`
    Duration     int       `json:"duration"`
    Source       string    `json:"source"`
}

type CassandraClient struct {
    session *gocql.Session
}

func NewCassandraClient() (*CassandraClient, error) {
    contactPoints := getEnv("CASSANDRA_CONTACT_POINTS", "localhost")
    keyspace := getEnv("CASSANDRA_KEYSPACE", "crm_keyspace")

    cluster := gocql.NewCluster(contactPoints)
    cluster.Keyspace = keyspace
    cluster.Consistency = gocql.Quorum
    cluster.ProtoVersion = 4
    cluster.ConnectTimeout = time.Second * 10
    cluster.Timeout = time.Second * 10

    session, err := cluster.CreateSession()
    if err != nil {
        return nil, fmt.Errorf("failed to connect to Cassandra: %w", err)
    }

    return &CassandraClient{session: session}, nil
}

func (c *CassandraClient) Close() {
    if c.session != nil {
        c.session.Close()
    }
}

func (c *CassandraClient) InsertActivity(activity Activity) error {
    query := `INSERT INTO customer_event_logs
              (customer_id, event_time, event_id, event_type, description, user_id, metadata, duration, source)
              VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`

    return c.session.Query(query,
        activity.CustomerID,
        activity.ActivityTime,
        gocql.TimeUUID(),
        activity.ActivityType,
        activity.Description,
        activity.UserID,
        activity.Metadata,
        activity.Duration,
        activity.Source,
    ).Exec()
}

func (c *CassandraClient) GetCustomerActivities(customerID int64, limit int) ([]Activity, error) {
    query := `SELECT customer_id, event_time, event_id, event_type, description, user_id, metadata, duration, source
              FROM customer_event_logs
              WHERE customer_id = ?
              ORDER BY event_time DESC
              LIMIT ?`

    iter := c.session.Query(query, customerID, limit).Iter()
    defer iter.Close()

    var activities []Activity
    var activity Activity
    var eventID gocql.UUID

    for iter.Scan(&activity.CustomerID, &activity.ActivityTime, &eventID, &activity.ActivityType,
        &activity.Description, &activity.UserID, &activity.Metadata, &activity.Duration, &activity.Source) {
        activity.ActivityID = eventID.String()
        activities = append(activities, activity)
    }

    if err := iter.Close(); err != nil {
        return nil, err
    }

    return activities, nil
}

func (c *CassandraClient) GetActivitiesByTimeRange(customerID int64, start, end time.Time) ([]Activity, error) {
    query := `SELECT customer_id, event_time, event_id, event_type, description, user_id, metadata, duration, source
              FROM customer_event_logs
              WHERE customer_id = ? AND event_time >= ? AND event_time <= ?
              ORDER BY event_time DESC`

    iter := c.session.Query(query, customerID, start, end).Iter()
    defer iter.Close()

    var activities []Activity
    var activity Activity
    var eventID gocql.UUID

    for iter.Scan(&activity.CustomerID, &activity.ActivityTime, &eventID, &activity.ActivityType,
        &activity.Description, &activity.UserID, &activity.Metadata, &activity.Duration, &activity.Source) {
        activity.ActivityID = eventID.String()
        activities = append(activities, activity)
    }

    if err := iter.Close(); err != nil {
        return nil, err
    }

    return activities, nil
}

func main() {
    // Initialize Cassandra
    cassandra, err := NewCassandraClient()
    if err != nil {
        log.Fatalf("Failed to connect to Cassandra: %v", err)
    }
    defer cassandra.Close()

    // Initialize Gin
    router := gin.Default()

    // Health check
    router.GET("/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "status":    "healthy",
            "service":   "activity-service",
            "timestamp": time.Now(),
        })
    })

    router.GET("/ready", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"status": "ready"})
    })

    // API routes
    api := router.Group("/api/activities")
    {
        api.POST("", func(c *gin.Context) {
            var activity Activity
            if err := c.ShouldBindJSON(&activity); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
                return
            }

            activity.ActivityTime = time.Now()
            if err := cassandra.InsertActivity(activity); err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }

            c.JSON(http.StatusCreated, activity)
        })

        api.GET("/customer/:id", func(c *gin.Context) {
            var customerID int64
            if _, err := fmt.Sscanf(c.Param("id"), "%d", &customerID); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
                return
            }

            limit := 100
            if limitParam := c.Query("limit"); limitParam != "" {
                fmt.Sscanf(limitParam, "%d", &limit)
            }

            activities, err := cassandra.GetCustomerActivities(customerID, limit)
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }

            c.JSON(http.StatusOK, activities)
        })

        api.GET("/customer/:id/range", func(c *gin.Context) {
            var customerID int64
            if _, err := fmt.Sscanf(c.Param("id"), "%d", &customerID); err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID"})
                return
            }

            startStr := c.Query("start")
            endStr := c.Query("end")

            start, err := time.Parse(time.RFC3339, startStr)
            if err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start time"})
                return
            }

            end, err := time.Parse(time.RFC3339, endStr)
            if err != nil {
                c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end time"})
                return
            }

            activities, err := cassandra.GetActivitiesByTimeRange(customerID, start, end)
            if err != nil {
                c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
                return
            }

            c.JSON(http.StatusOK, activities)
        })
    }

    // Start server
    srv := &http.Server{
        Addr:    ":5007",
        Handler: router,
    }

    go func() {
        log.Println("Activity Service starting on port 5007...")
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("Failed to start server: %v", err)
        }
    }()

    // Graceful shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit

    log.Println("Shutting down server...")
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("Server forced to shutdown:", err)
    }

    log.Println("Server exited")
}

func getEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}

package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server
	ServerPort  string
	Environment string

	// AI Provider
	AIProvider     string
	ClaudeAPIKey   string
	OpenAIAPIKey   string
	ClaudeModel    string
	OpenAIModel    string

	// Database
	PostgresHost     string
	PostgresPort     string
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string

	// Redis
	RedisHost     string
	RedisPort     string
	RedisPassword string

	// RabbitMQ
	RabbitMQHost     string
	RabbitMQPort     string
	RabbitMQUser     string
	RabbitMQPassword string
	RabbitMQVHost    string

	// Kafka
	KafkaBrokers string
	KafkaGroupID string

	// WebSocket
	WebSocketPath string
}

func LoadConfig() (*Config, error) {
	// Load .env file if exists
	_ = godotenv.Load()

	return &Config{
		ServerPort:  getEnv("CHAT_SERVICE_PORT", "8082"),
		Environment: getEnv("ENVIRONMENT", "development"),

		AIProvider:     getEnv("AI_PROVIDER", "claude"),
		ClaudeAPIKey:   getEnv("CLAUDE_API_KEY", ""),
		OpenAIAPIKey:   getEnv("OPENAI_API_KEY", ""),
		ClaudeModel:    getEnv("AI_MODEL_CLAUDE", "claude-3-5-sonnet-20241022"),
		OpenAIModel:    getEnv("AI_MODEL_OPENAI", "gpt-4-turbo-preview"),

		PostgresHost:     getEnv("POSTGRES_HOST", "localhost"),
		PostgresPort:     getEnv("POSTGRES_PORT", "5432"),
		PostgresDB:       getEnv("POSTGRES_DB", "crm_db"),
		PostgresUser:     getEnv("POSTGRES_USER", "crm_user"),
		PostgresPassword: getEnv("POSTGRES_PASSWORD", "password"),

		RedisHost:     getEnv("REDIS_HOST", "localhost"),
		RedisPort:     getEnv("REDIS_PORT", "6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),

		RabbitMQHost:     getEnv("RABBITMQ_HOST", "localhost"),
		RabbitMQPort:     getEnv("RABBITMQ_PORT", "5672"),
		RabbitMQUser:     getEnv("RABBITMQ_USER", "crm_user"),
		RabbitMQPassword: getEnv("RABBITMQ_PASSWORD", "password"),
		RabbitMQVHost:    getEnv("RABBITMQ_VHOST", "/crm"),

		KafkaBrokers: getEnv("KAFKA_BROKERS", "localhost:9092"),
		KafkaGroupID: getEnv("KAFKA_GROUP_ID", "crm-services"),

		WebSocketPath: getEnv("WEBSOCKET_PATH", "/ws"),
	}, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

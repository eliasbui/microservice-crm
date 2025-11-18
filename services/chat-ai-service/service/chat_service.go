package service

import (
	"context"
	"fmt"
	"time"

	"github.com/crm/chat-ai-service/config"
	"github.com/crm/chat-ai-service/models"
	"github.com/google/uuid"
)

type ChatService struct {
	config    *config.Config
	aiService *AIService
}

func NewChatService(cfg *config.Config, aiService *AIService) (*ChatService, error) {
	return &ChatService{
		config:    cfg,
		aiService: aiService,
	}, nil
}

func (s *ChatService) ProcessMessage(ctx context.Context, req models.ChatRequest) (*models.ChatResponse, error) {
	// Save user message
	userMessage := models.Message{
		ID:        uuid.New().String(),
		SessionID: req.SessionID,
		UserID:    req.UserID,
		Content:   req.Message,
		Role:      "user",
		Timestamp: time.Now(),
	}

	// Get AI response
	aiResponse, err := s.aiService.Chat(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to get AI response: %w", err)
	}

	// Save assistant message
	assistantMessage := models.Message{
		ID:        uuid.New().String(),
		SessionID: req.SessionID,
		UserID:    "assistant",
		Content:   aiResponse.Message,
		Role:      "assistant",
		Timestamp: time.Now(),
		Metadata:  aiResponse.Metadata,
	}

	// TODO: Persist messages to database

	_ = userMessage
	_ = assistantMessage

	return aiResponse, nil
}

func (s *ChatService) GetMessageHistory(sessionID string) ([]models.Message, error) {
	// TODO: Retrieve message history from database
	return []models.Message{}, nil
}

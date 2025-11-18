package service

import (
	"context"
	"fmt"

	"github.com/crm/chat-ai-service/config"
	"github.com/crm/chat-ai-service/models"
	"github.com/sashabaranov/go-openai"
)

type AIService struct {
	config       *config.Config
	openaiClient *openai.Client
}

func NewAIService(cfg *config.Config) (*AIService, error) {
	var openaiClient *openai.Client

	if cfg.AIProvider == "openai" && cfg.OpenAIAPIKey != "" {
		openaiClient = openai.NewClient(cfg.OpenAIAPIKey)
	}

	return &AIService{
		config:       cfg,
		openaiClient: openaiClient,
	}, nil
}

func (s *AIService) Chat(ctx context.Context, req models.ChatRequest) (*models.ChatResponse, error) {
	if s.config.AIProvider == "claude" {
		return s.chatWithClaude(ctx, req)
	} else if s.config.AIProvider == "openai" {
		return s.chatWithOpenAI(ctx, req)
	}

	return nil, fmt.Errorf("unsupported AI provider: %s", s.config.AIProvider)
}

func (s *AIService) chatWithClaude(ctx context.Context, req models.ChatRequest) (*models.ChatResponse, error) {
	// Implementation for Claude API
	// Note: This is a placeholder - you would use the actual Claude SDK here
	return &models.ChatResponse{
		SessionID: req.SessionID,
		Message:   "This is a response from Claude AI (placeholder implementation)",
		Metadata: models.Metadata{
			Model:      s.config.ClaudeModel,
			Tokens:     100,
			Confidence: 0.95,
		},
	}, nil
}

func (s *AIService) chatWithOpenAI(ctx context.Context, req models.ChatRequest) (*models.ChatResponse, error) {
	if s.openaiClient == nil {
		return nil, fmt.Errorf("OpenAI client not initialized")
	}

	// Convert messages to OpenAI format
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: "You are a helpful CRM assistant. Help users manage customers and opportunities.",
		},
	}

	// Add context messages
	for _, msg := range req.Context {
		role := openai.ChatMessageRoleUser
		if msg.Role == "assistant" {
			role = openai.ChatMessageRoleAssistant
		}
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    role,
			Content: msg.Content,
		})
	}

	// Add current message
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: req.Message,
	})

	resp, err := s.openaiClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:    s.config.OpenAIModel,
		Messages: messages,
	})

	if err != nil {
		return nil, fmt.Errorf("OpenAI API error: %w", err)
	}

	return &models.ChatResponse{
		SessionID: req.SessionID,
		Message:   resp.Choices[0].Message.Content,
		Metadata: models.Metadata{
			Model:  s.config.OpenAIModel,
			Tokens: resp.Usage.TotalTokens,
		},
	}, nil
}

func (s *AIService) AnalyzeSentiment(ctx context.Context, text string) (*models.SentimentResponse, error) {
	// Simplified sentiment analysis using AI
	prompt := fmt.Sprintf("Analyze the sentiment of the following text and respond with only: positive, negative, or neutral.\n\nText: %s", text)

	req := models.ChatRequest{
		SessionID: "sentiment-analysis",
		Message:   prompt,
	}

	resp, err := s.Chat(ctx, req)
	if err != nil {
		return nil, err
	}

	return &models.SentimentResponse{
		Sentiment:  resp.Message,
		Confidence: resp.Metadata.Confidence,
	}, nil
}

func (s *AIService) ClassifyIntent(ctx context.Context, text string) (*models.IntentResponse, error) {
	// Simplified intent classification
	prompt := fmt.Sprintf("Classify the intent of this customer message. Choose from: inquiry, complaint, request, feedback, other.\n\nMessage: %s", text)

	req := models.ChatRequest{
		SessionID: "intent-classification",
		Message:   prompt,
	}

	resp, err := s.Chat(ctx, req)
	if err != nil {
		return nil, err
	}

	return &models.IntentResponse{
		Intent:     resp.Message,
		Confidence: resp.Metadata.Confidence,
	}, nil
}

func (s *AIService) GenerateEmail(ctx context.Context, context, tone, purpose string) (*models.EmailGenerationResponse, error) {
	prompt := fmt.Sprintf("Generate a professional email with the following details:\nContext: %s\nTone: %s\nPurpose: %s\n\nProvide the subject line and body separately.",
		context, tone, purpose)

	req := models.ChatRequest{
		SessionID: "email-generation",
		Message:   prompt,
	}

	resp, err := s.Chat(ctx, req)
	if err != nil {
		return nil, err
	}

	return &models.EmailGenerationResponse{
		Subject: "Generated Email Subject",
		Body:    resp.Message,
	}, nil
}

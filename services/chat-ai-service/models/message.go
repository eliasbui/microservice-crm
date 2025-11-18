package models

import "time"

type Message struct {
	ID        string    `json:"id"`
	SessionID string    `json:"session_id"`
	UserID    string    `json:"user_id"`
	Content   string    `json:"content"`
	Role      string    `json:"role"` // user, assistant, system
	Timestamp time.Time `json:"timestamp"`
	Metadata  Metadata  `json:"metadata,omitempty"`
}

type Metadata struct {
	Model      string  `json:"model,omitempty"`
	Tokens     int     `json:"tokens,omitempty"`
	Confidence float64 `json:"confidence,omitempty"`
}

type ChatRequest struct {
	SessionID string    `json:"session_id"`
	UserID    string    `json:"user_id"`
	Message   string    `json:"message"`
	Context   []Message `json:"context,omitempty"`
}

type ChatResponse struct {
	SessionID string    `json:"session_id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
	Metadata  Metadata  `json:"metadata,omitempty"`
}

type SentimentRequest struct {
	Text string `json:"text" binding:"required"`
}

type SentimentResponse struct {
	Sentiment  string  `json:"sentiment"` // positive, negative, neutral
	Confidence float64 `json:"confidence"`
	Details    string  `json:"details,omitempty"`
}

type IntentRequest struct {
	Text string `json:"text" binding:"required"`
}

type IntentResponse struct {
	Intent     string  `json:"intent"`
	Confidence float64 `json:"confidence"`
	Entities   []Entity `json:"entities,omitempty"`
}

type Entity struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type EmailGenerationRequest struct {
	Context string `json:"context" binding:"required"`
	Tone    string `json:"tone,omitempty"` // formal, friendly, professional
	Purpose string `json:"purpose,omitempty"`
}

type EmailGenerationResponse struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

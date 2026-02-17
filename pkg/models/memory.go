package models

import (
	"time"
)

// MemoryType represents the category of a memory
type MemoryType string

const (
	MemoryTypeDecision   MemoryType = "decision"
	MemoryTypePattern    MemoryType = "pattern"
	MemoryTypeFact       MemoryType = "fact"
	MemoryTypePreference MemoryType = "preference"
	MemoryTypeMistake    MemoryType = "mistake"
	MemoryTypeLearning   MemoryType = "learning"
)

// MemoryScope represents the visibility of a memory
type MemoryScope string

const (
	MemoryScopePersonal MemoryScope = "personal"
	MemoryScopeProject  MemoryScope = "project"
	MemoryScopeTeam     MemoryScope = "team"
	MemoryScopeOrg      MemoryScope = "org"
)

// SourceType represents where a memory came from
type SourceType string

const (
	SourceTypeGit      SourceType = "git"
	SourceTypeFile     SourceType = "file"
	SourceTypeTerminal SourceType = "terminal"
	SourceTypeChat     SourceType = "chat"
	SourceTypeManual   SourceType = "manual"
	SourceTypeImport   SourceType = "import"
)

// Source tracks where a memory originated
type Source struct {
	Type      SourceType `json:"type"`
	Reference string     `json:"reference"` // commit hash, file path, etc.
	Timestamp time.Time  `json:"timestamp"`
}

// Memory represents a single piece of remembered information
type Memory struct {
	ID      string     `json:"id"`
	Type    MemoryType `json:"type"`
	Content string     `json:"content"`
	Summary string     `json:"summary"`

	// Scope
	Scope     MemoryScope `json:"scope"`
	ProjectID *string     `json:"projectId,omitempty"`
	TeamID    *string     `json:"teamId,omitempty"`

	// Source tracking
	Source Source `json:"source"`

	// Intelligence
	Confidence float64   `json:"confidence"` // 0.0-1.0
	Importance float64   `json:"importance"` // 0.0-1.0, decays over time
	Embedding  []float32 `json:"embedding"`  // 384-dim vector

	// Relationships
	Topics          []string `json:"topics"`
	RelatedMemories []string `json:"relatedMemories"`

	// Lifecycle
	CreatedAt      time.Time  `json:"createdAt"`
	LastAccessedAt time.Time  `json:"lastAccessedAt"`
	AccessCount    int        `json:"accessCount"`
	ExpiresAt      *time.Time `json:"expiresAt,omitempty"`
}

// Project represents a tracked project/repository
type Project struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	GitRemote *string   `json:"gitRemote,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	LastSeen  time.Time `json:"lastSeen"`
}

// Event represents a captured event from watchers
type Event struct {
	ID        string                 `json:"id"`
	Type      string                 `json:"type"` // git_commit, file_change, terminal_cmd
	Timestamp time.Time              `json:"timestamp"`
	Data      map[string]interface{} `json:"data"`
	ProjectID *string                `json:"projectId,omitempty"`
}

// RecallRequest represents a search query
type RecallRequest struct {
	Query     string        `json:"query"`
	Scope     []MemoryScope `json:"scope,omitempty"`
	ProjectID *string       `json:"projectId,omitempty"`
	Types     []MemoryType  `json:"types,omitempty"`
	Limit     int           `json:"limit,omitempty"`
}

// RecallResponse represents search results
type RecallResponse struct {
	Memories []Memory `json:"memories"`
	Total    int      `json:"total"`
	Query    string   `json:"query"`
}

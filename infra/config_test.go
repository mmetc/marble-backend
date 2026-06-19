package infra

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/genai"
)

func TestNewAIAgentConfigurationDefaults(t *testing.T) {
	config := NewAIAgentConfiguration("gcp-project")

	assert.Equal(t, AIAgentProviderTypeOpenAI, config.MainAgentProviderType)
	assert.Equal(t, "gemini-2.5-flash", config.MainAgentDefaultModel)
	assert.Equal(t, "gcp-project", config.MainAgentProject)
}

func TestNewAIAgentConfigurationReadsEnv(t *testing.T) {
	t.Setenv("AI_AGENT_MAIN_AGENT_PROVIDER_TYPE", "aistudio")
	t.Setenv("AI_AGENT_MAIN_AGENT_URL", "https://example.test")
	t.Setenv("AI_AGENT_MAIN_AGENT_KEY", "secret")
	t.Setenv("AI_AGENT_MAIN_AGENT_DEFAULT_MODEL", "gpt-test")
	t.Setenv("AI_AGENT_MAIN_AGENT_BACKEND", "vertex")
	t.Setenv("AI_AGENT_MAIN_AGENT_PROJECT", "explicit-project")
	t.Setenv("AI_AGENT_MAIN_AGENT_LOCATION", "europe-west1")
	t.Setenv("AI_AGENT_PERPLEXITY_API_KEY", "pplx-key")

	config := NewAIAgentConfiguration("fallback-project")

	assert.Equal(t, AIAgentProviderTypeAIStudio, config.MainAgentProviderType)
	assert.Equal(t, "https://example.test", config.MainAgentURL)
	assert.Equal(t, "secret", config.MainAgentKey)
	assert.Equal(t, "gpt-test", config.MainAgentDefaultModel)
	assert.Equal(t, genai.BackendVertexAI, config.MainAgentBackend)
	assert.Equal(t, "explicit-project", config.MainAgentProject)
	assert.Equal(t, "europe-west1", config.MainAgentLocation)
	assert.Equal(t, "pplx-key", config.PerplexityAPIKey)
}

package agent

import "fmt"

type AgentModel string

const (
	AgentModelGPT3Dot5                AgentModel = "gpt-3.5"
	AgentModelGPT4                    AgentModel = "gpt-4"
	AgentModel3Mini                   AgentModel = "o3-mini"
	AgentModel4Mini                   AgentModel = "o4-mini"
	AgentModelClaudeSonnet3Dot5       AgentModel = "claude-sonnet-3.5"
	AgentModelClaudeSonnet3Dot7       AgentModel = "claude-sonnet-3.7"
	AgentModelDeepSeekR1              AgentModel = "deepseek-reasoner"
	AgentModelGeminiFlash1Dot5        AgentModel = "gemini-1.5-flash"
	AgentModelGeminiFlash2Dot0        AgentModel = "gemini-2.0-flash"
	AgentModelGemini2Dot5ProPreview   AgentModel = "gemini-2.5-pro-preview-03-25"
	AgentModelGemini2Dot5FlashPreview AgentModel = "gemini-2.5-flash-preview-04-17"
	AgentModelGrok2Dot0               AgentModel = "grok-2-latest"
	AgentModelNone                    AgentModel = "none"
)

type ModeCatalog struct {
	Name  string     `json:"name"`
	Model AgentModel `json:"model"`
}

var AvailableModelCatalogs = []ModeCatalog{
	{
		Name:  "GPT 3.5",
		Model: AgentModelGPT3Dot5,
	},
	{
		Name:  "GPT 4",
		Model: AgentModelGPT4,
	},
	{
		Name:  "o3 Mini",
		Model: AgentModel3Mini,
	},
	{
		Name:  "o4 Mini",
		Model: AgentModel4Mini,
	},
	{
		Name:  "Claude Sonnet 3.5",
		Model: AgentModelClaudeSonnet3Dot5,
	},
	{
		Name:  "Claude Sonnet 3.7",
		Model: AgentModelClaudeSonnet3Dot7,
	},
	{
		Name:  "DeepSeek R1",
		Model: AgentModelDeepSeekR1,
	},
	{
		Name:  "Gemini 1.5 Flash",
		Model: AgentModelGeminiFlash1Dot5,
	},
	{
		Name:  "Gemini 2.0 Flash",
		Model: AgentModelGeminiFlash2Dot0,
	},
	{
		Name:  "Gemini 2.5 Pro Preview",
		Model: AgentModelGemini2Dot5ProPreview,
	},
	{
		Name:  "Gemini 2.5 Flash Preview",
		Model: AgentModelGemini2Dot5FlashPreview,
	},
	{
		Name:  "Grok 2.0",
		Model: AgentModelGrok2Dot0,
	},
}

func GetModelCatalogByIndex(index string) (ModeCatalog, error) {
	for i, catalog := range AvailableModelCatalogs {
		if fmt.Sprint(i) == index {
			return catalog, nil
		}
	}
	return ModeCatalog{}, fmt.Errorf("model not found")
}

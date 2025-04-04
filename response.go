package openai

type Response struct {
	ID                 string                 `json:"id"`
	Object             string                 `json:"object"`
	CreatedAt          int64                  `json:"created_at"`
	Status             string                 `json:"status"`
	Error              interface{}            `json:"error"`
	IncompleteDetails  interface{}            `json:"incomplete_details"`
	Instructions       interface{}            `json:"instructions"`
	MaxOutputTokens    interface{}            `json:"max_output_tokens"`
	Model              string                 `json:"model"`
	Output             []OutputItem           `json:"output"`
	ParallelToolCalls  bool                   `json:"parallel_tool_calls"`
	PreviousResponseID interface{}            `json:"previous_response_id"`
	Reasoning          Reasoning              `json:"reasoning"`
	Store              bool                   `json:"store"`
	Temperature        float64                `json:"temperature"`
	Text               TextFormat             `json:"text"`
	ToolChoice         string                 `json:"tool_choice"`
	Tools              []interface{}          `json:"tools"`
	TopP               float64                `json:"top_p"`
	Truncation         string                 `json:"truncation"`
	Usage              Usage                  `json:"usage"`
	User               interface{}            `json:"user"`
	Metadata           map[string]interface{} `json:"metadata"`
}

type OutputItem struct {
	Type    string         `json:"type"`
	ID      string         `json:"id"`
	Status  string         `json:"status"`
	Role    string         `json:"role"`
	Content []ContentBlock `json:"content"`
}

type ContentBlock struct {
	Type        string        `json:"type"`
	Text        string        `json:"text"`
	Annotations []interface{} `json:"annotations"`
}

type Reasoning struct {
	Effort          interface{} `json:"effort"`
	GenerateSummary interface{} `json:"generate_summary"`
}

type TextFormat struct {
	Format FormatType `json:"format"`
}

type FormatType struct {
	Type string `json:"type"`
}

type Usage struct {
	InputTokens         int                `json:"input_tokens"`
	InputTokensDetails  TokenDetails       `json:"input_tokens_details"`
	OutputTokens        int                `json:"output_tokens"`
	OutputTokensDetails OutputTokenDetails `json:"output_tokens_details"`
	TotalTokens         int                `json:"total_tokens"`
}

type TokenDetails struct {
	CachedTokens int `json:"cached_tokens"`
}

type OutputTokenDetails struct {
	ReasoningTokens int `json:"reasoning_tokens"`
}

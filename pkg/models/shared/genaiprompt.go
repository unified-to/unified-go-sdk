// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GenaiPrompt struct {
	MaxTokens   *float64       `json:"max_tokens,omitempty"`
	Messages    []GenaiContent `json:"messages,omitempty"`
	ModelID     *string        `json:"model_id,omitempty"`
	Raw         map[string]any `json:"raw,omitempty"`
	Responses   []string       `json:"responses,omitempty"`
	Temperature *float64       `json:"temperature,omitempty"`
}

func (o *GenaiPrompt) GetMaxTokens() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxTokens
}

func (o *GenaiPrompt) GetMessages() []GenaiContent {
	if o == nil {
		return nil
	}
	return o.Messages
}

func (o *GenaiPrompt) GetModelID() *string {
	if o == nil {
		return nil
	}
	return o.ModelID
}

func (o *GenaiPrompt) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *GenaiPrompt) GetResponses() []string {
	if o == nil {
		return nil
	}
	return o.Responses
}

func (o *GenaiPrompt) GetTemperature() *float64 {
	if o == nil {
		return nil
	}
	return o.Temperature
}
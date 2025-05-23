// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type AtsJobQuestionType string

const (
	AtsJobQuestionTypeText           AtsJobQuestionType = "TEXT"
	AtsJobQuestionTypeNumber         AtsJobQuestionType = "NUMBER"
	AtsJobQuestionTypeDate           AtsJobQuestionType = "DATE"
	AtsJobQuestionTypeBoolean        AtsJobQuestionType = "BOOLEAN"
	AtsJobQuestionTypeMultipleChoice AtsJobQuestionType = "MULTIPLE_CHOICE"
	AtsJobQuestionTypeFile           AtsJobQuestionType = "FILE"
	AtsJobQuestionTypeTextarea       AtsJobQuestionType = "TEXTAREA"
	AtsJobQuestionTypeMultipleSelect AtsJobQuestionType = "MULTIPLE_SELECT"
	AtsJobQuestionTypeUniversity     AtsJobQuestionType = "UNIVERSITY"
	AtsJobQuestionTypeYesNo          AtsJobQuestionType = "YES_NO"
	AtsJobQuestionTypeCurrency       AtsJobQuestionType = "CURRENCY"
	AtsJobQuestionTypeURL            AtsJobQuestionType = "URL"
)

func (e AtsJobQuestionType) ToPointer() *AtsJobQuestionType {
	return &e
}
func (e *AtsJobQuestionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TEXT":
		fallthrough
	case "NUMBER":
		fallthrough
	case "DATE":
		fallthrough
	case "BOOLEAN":
		fallthrough
	case "MULTIPLE_CHOICE":
		fallthrough
	case "FILE":
		fallthrough
	case "TEXTAREA":
		fallthrough
	case "MULTIPLE_SELECT":
		fallthrough
	case "UNIVERSITY":
		fallthrough
	case "YES_NO":
		fallthrough
	case "CURRENCY":
		fallthrough
	case "URL":
		*e = AtsJobQuestionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AtsJobQuestionType: %v", v)
	}
}

type AtsJobQuestion struct {
	Description *string                         `json:"description,omitempty"`
	ID          *string                         `json:"id,omitempty"`
	Options     []PropertyAtsJobQuestionOptions `json:"options,omitempty"`
	Prompt      *string                         `json:"prompt,omitempty"`
	Question    string                          `json:"question"`
	Required    *bool                           `json:"required,omitempty"`
	Type        AtsJobQuestionType              `json:"type"`
}

func (o *AtsJobQuestion) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AtsJobQuestion) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AtsJobQuestion) GetOptions() []PropertyAtsJobQuestionOptions {
	if o == nil {
		return nil
	}
	return o.Options
}

func (o *AtsJobQuestion) GetPrompt() *string {
	if o == nil {
		return nil
	}
	return o.Prompt
}

func (o *AtsJobQuestion) GetQuestion() string {
	if o == nil {
		return ""
	}
	return o.Question
}

func (o *AtsJobQuestion) GetRequired() *bool {
	if o == nil {
		return nil
	}
	return o.Required
}

func (o *AtsJobQuestion) GetType() AtsJobQuestionType {
	if o == nil {
		return AtsJobQuestionType("")
	}
	return o.Type
}

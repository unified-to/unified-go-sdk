// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type AtsActivityType string

const (
	AtsActivityTypeNote  AtsActivityType = "NOTE"
	AtsActivityTypeTask  AtsActivityType = "TASK"
	AtsActivityTypeEmail AtsActivityType = "EMAIL"
)

func (e AtsActivityType) ToPointer() *AtsActivityType {
	return &e
}
func (e *AtsActivityType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NOTE":
		fallthrough
	case "TASK":
		fallthrough
	case "EMAIL":
		*e = AtsActivityType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AtsActivityType: %v", v)
	}
}

type AtsActivity struct {
	ApplicationID *string                  `json:"application_id,omitempty"`
	Bcc           []AtsEmail               `json:"bcc,omitempty"`
	CandidateID   *string                  `json:"candidate_id,omitempty"`
	Cc            []AtsEmail               `json:"cc,omitempty"`
	CreatedAt     *time.Time               `json:"created_at,omitempty"`
	Description   *string                  `json:"description,omitempty"`
	DocumentID    *string                  `json:"document_id,omitempty"`
	From          *PropertyAtsActivityFrom `json:"from,omitempty"`
	ID            *string                  `json:"id,omitempty"`
	InterviewID   *string                  `json:"interview_id,omitempty"`
	IsPrivate     *bool                    `json:"is_private,omitempty"`
	JobID         *string                  `json:"job_id,omitempty"`
	Raw           map[string]any           `json:"raw,omitempty"`
	SubType       *string                  `json:"sub_type,omitempty"`
	Title         *string                  `json:"title,omitempty"`
	To            []AtsEmail               `json:"to,omitempty"`
	Type          *AtsActivityType         `json:"type,omitempty"`
	UpdatedAt     *time.Time               `json:"updated_at,omitempty"`
	// id values of the recruiters associated with the activity.
	UserIds []string `json:"user_ids,omitempty"`
}

func (a AtsActivity) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AtsActivity) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AtsActivity) GetApplicationID() *string {
	if o == nil {
		return nil
	}
	return o.ApplicationID
}

func (o *AtsActivity) GetBcc() []AtsEmail {
	if o == nil {
		return nil
	}
	return o.Bcc
}

func (o *AtsActivity) GetCandidateID() *string {
	if o == nil {
		return nil
	}
	return o.CandidateID
}

func (o *AtsActivity) GetCc() []AtsEmail {
	if o == nil {
		return nil
	}
	return o.Cc
}

func (o *AtsActivity) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AtsActivity) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AtsActivity) GetDocumentID() *string {
	if o == nil {
		return nil
	}
	return o.DocumentID
}

func (o *AtsActivity) GetFrom() *PropertyAtsActivityFrom {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *AtsActivity) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AtsActivity) GetInterviewID() *string {
	if o == nil {
		return nil
	}
	return o.InterviewID
}

func (o *AtsActivity) GetIsPrivate() *bool {
	if o == nil {
		return nil
	}
	return o.IsPrivate
}

func (o *AtsActivity) GetJobID() *string {
	if o == nil {
		return nil
	}
	return o.JobID
}

func (o *AtsActivity) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *AtsActivity) GetSubType() *string {
	if o == nil {
		return nil
	}
	return o.SubType
}

func (o *AtsActivity) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *AtsActivity) GetTo() []AtsEmail {
	if o == nil {
		return nil
	}
	return o.To
}

func (o *AtsActivity) GetType() *AtsActivityType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *AtsActivity) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *AtsActivity) GetUserIds() []string {
	if o == nil {
		return nil
	}
	return o.UserIds
}

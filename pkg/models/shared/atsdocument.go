// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type AtsDocumentType string

const (
	AtsDocumentTypeResume       AtsDocumentType = "RESUME"
	AtsDocumentTypeCoverLetter  AtsDocumentType = "COVER_LETTER"
	AtsDocumentTypeOfferPacket  AtsDocumentType = "OFFER_PACKET"
	AtsDocumentTypeOfferLetter  AtsDocumentType = "OFFER_LETTER"
	AtsDocumentTypeTakeHomeTest AtsDocumentType = "TAKE_HOME_TEST"
	AtsDocumentTypeOther        AtsDocumentType = "OTHER"
)

func (e AtsDocumentType) ToPointer() *AtsDocumentType {
	return &e
}

func (e *AtsDocumentType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "RESUME":
		fallthrough
	case "COVER_LETTER":
		fallthrough
	case "OFFER_PACKET":
		fallthrough
	case "OFFER_LETTER":
		fallthrough
	case "TAKE_HOME_TEST":
		fallthrough
	case "OTHER":
		*e = AtsDocumentType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AtsDocumentType: %v", v)
	}
}

type AtsDocument struct {
	ApplicationID *string                `json:"application_id,omitempty"`
	CandidateID   *string                `json:"candidate_id,omitempty"`
	CreatedAt     *time.Time             `json:"created_at,omitempty"`
	DocumentData  *string                `json:"document_data,omitempty"`
	DocumentURL   *string                `json:"document_url,omitempty"`
	Filename      *string                `json:"filename,omitempty"`
	ID            *string                `json:"id,omitempty"`
	JobID         *string                `json:"job_id,omitempty"`
	Raw           map[string]interface{} `json:"raw,omitempty"`
	Type          *AtsDocumentType       `json:"type,omitempty"`
	UpdatedAt     *time.Time             `json:"updated_at,omitempty"`
	UserID        *string                `json:"user_id,omitempty"`
}

func (a AtsDocument) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AtsDocument) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AtsDocument) GetApplicationID() *string {
	if o == nil {
		return nil
	}
	return o.ApplicationID
}

func (o *AtsDocument) GetCandidateID() *string {
	if o == nil {
		return nil
	}
	return o.CandidateID
}

func (o *AtsDocument) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AtsDocument) GetDocumentData() *string {
	if o == nil {
		return nil
	}
	return o.DocumentData
}

func (o *AtsDocument) GetDocumentURL() *string {
	if o == nil {
		return nil
	}
	return o.DocumentURL
}

func (o *AtsDocument) GetFilename() *string {
	if o == nil {
		return nil
	}
	return o.Filename
}

func (o *AtsDocument) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AtsDocument) GetJobID() *string {
	if o == nil {
		return nil
	}
	return o.JobID
}

func (o *AtsDocument) GetRaw() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *AtsDocument) GetType() *AtsDocumentType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *AtsDocument) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *AtsDocument) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

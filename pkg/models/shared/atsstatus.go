// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type AtsStatusStatus string

const (
	AtsStatusStatusNew             AtsStatusStatus = "NEW"
	AtsStatusStatusReviewing       AtsStatusStatus = "REVIEWING"
	AtsStatusStatusScreening       AtsStatusStatus = "SCREENING"
	AtsStatusStatusSubmitted       AtsStatusStatus = "SUBMITTED"
	AtsStatusStatusFirstInterview  AtsStatusStatus = "FIRST_INTERVIEW"
	AtsStatusStatusSecondInterview AtsStatusStatus = "SECOND_INTERVIEW"
	AtsStatusStatusThirdInterview  AtsStatusStatus = "THIRD_INTERVIEW"
	AtsStatusStatusBackgroundCheck AtsStatusStatus = "BACKGROUND_CHECK"
	AtsStatusStatusOffered         AtsStatusStatus = "OFFERED"
	AtsStatusStatusAccepted        AtsStatusStatus = "ACCEPTED"
	AtsStatusStatusHired           AtsStatusStatus = "HIRED"
	AtsStatusStatusRejected        AtsStatusStatus = "REJECTED"
	AtsStatusStatusWithdrawn       AtsStatusStatus = "WITHDRAWN"
)

func (e AtsStatusStatus) ToPointer() *AtsStatusStatus {
	return &e
}

func (e *AtsStatusStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NEW":
		fallthrough
	case "REVIEWING":
		fallthrough
	case "SCREENING":
		fallthrough
	case "SUBMITTED":
		fallthrough
	case "FIRST_INTERVIEW":
		fallthrough
	case "SECOND_INTERVIEW":
		fallthrough
	case "THIRD_INTERVIEW":
		fallthrough
	case "BACKGROUND_CHECK":
		fallthrough
	case "OFFERED":
		fallthrough
	case "ACCEPTED":
		fallthrough
	case "HIRED":
		fallthrough
	case "REJECTED":
		fallthrough
	case "WITHDRAWN":
		*e = AtsStatusStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AtsStatusStatus: %v", v)
	}
}

type AtsStatus struct {
	Description    *string          `json:"description,omitempty"`
	ID             *string          `json:"id,omitempty"`
	OriginalStatus *string          `json:"original_status,omitempty"`
	Raw            map[string]any   `json:"raw,omitempty"`
	Status         *AtsStatusStatus `json:"status,omitempty"`
}

func (o *AtsStatus) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AtsStatus) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AtsStatus) GetOriginalStatus() *string {
	if o == nil {
		return nil
	}
	return o.OriginalStatus
}

func (o *AtsStatus) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *AtsStatus) GetStatus() *AtsStatusStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

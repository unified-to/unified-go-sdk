// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type CrmEventType string

const (
	CrmEventTypeNote    CrmEventType = "NOTE"
	CrmEventTypeEmail   CrmEventType = "EMAIL"
	CrmEventTypeTask    CrmEventType = "TASK"
	CrmEventTypeMeeting CrmEventType = "MEETING"
	CrmEventTypeCall    CrmEventType = "CALL"
)

func (e CrmEventType) ToPointer() *CrmEventType {
	return &e
}

func (e *CrmEventType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "NOTE":
		fallthrough
	case "EMAIL":
		fallthrough
	case "TASK":
		fallthrough
	case "MEETING":
		fallthrough
	case "CALL":
		*e = CrmEventType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CrmEventType: %v", v)
	}
}

// CrmEvent - An event represents an event, activity, or engagement and is always associated with a deal, contact, or company
type CrmEvent struct {
	// The call object, when type = call
	Call *PropertyCrmEventCall `json:"call,omitempty"`
	// An array of company IDs associated with this event
	CompanyIds []string `json:"company_ids,omitempty"`
	// An array of contact IDs associated with this event
	ContactIds []string   `json:"contact_ids,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	// An array of deal IDs associated with this event
	DealIds []string `json:"deal_ids,omitempty"`
	// The email object, when type = email
	Email   *PropertyCrmEventEmail `json:"email,omitempty"`
	ID      *string                `json:"id,omitempty"`
	LeadIds []string               `json:"lead_ids,omitempty"`
	// The meeting object, when type = meeting
	Meeting *PropertyCrmEventMeeting `json:"meeting,omitempty"`
	// The note object, when type = note
	Note *PropertyCrmEventNote `json:"note,omitempty"`
	// The raw data returned by the integration for this event.
	Raw *PropertyCrmEventRaw `json:"raw,omitempty"`
	// The task object, when type = task
	Task      *PropertyCrmEventTask `json:"task,omitempty"`
	Type      *CrmEventType         `json:"type,omitempty"`
	UpdatedAt *time.Time            `json:"updated_at,omitempty"`
}

func (c CrmEvent) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CrmEvent) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CrmEvent) GetCall() *PropertyCrmEventCall {
	if o == nil {
		return nil
	}
	return o.Call
}

func (o *CrmEvent) GetCompanyIds() []string {
	if o == nil {
		return nil
	}
	return o.CompanyIds
}

func (o *CrmEvent) GetContactIds() []string {
	if o == nil {
		return nil
	}
	return o.ContactIds
}

func (o *CrmEvent) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CrmEvent) GetDealIds() []string {
	if o == nil {
		return nil
	}
	return o.DealIds
}

func (o *CrmEvent) GetEmail() *PropertyCrmEventEmail {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *CrmEvent) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CrmEvent) GetLeadIds() []string {
	if o == nil {
		return nil
	}
	return o.LeadIds
}

func (o *CrmEvent) GetMeeting() *PropertyCrmEventMeeting {
	if o == nil {
		return nil
	}
	return o.Meeting
}

func (o *CrmEvent) GetNote() *PropertyCrmEventNote {
	if o == nil {
		return nil
	}
	return o.Note
}

func (o *CrmEvent) GetRaw() *PropertyCrmEventRaw {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *CrmEvent) GetTask() *PropertyCrmEventTask {
	if o == nil {
		return nil
	}
	return o.Task
}

func (o *CrmEvent) GetType() *CrmEventType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *CrmEvent) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

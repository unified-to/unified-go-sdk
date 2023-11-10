// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type TicketingTicketStatus string

const (
	TicketingTicketStatusActive TicketingTicketStatus = "ACTIVE"
	TicketingTicketStatusClosed TicketingTicketStatus = "CLOSED"
)

func (e TicketingTicketStatus) ToPointer() *TicketingTicketStatus {
	return &e
}

func (e *TicketingTicketStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ACTIVE":
		fallthrough
	case "CLOSED":
		*e = TicketingTicketStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TicketingTicketStatus: %v", v)
	}
}

type TicketingTicket struct {
	Category    *string                    `json:"category,omitempty"`
	ClosedAt    *time.Time                 `json:"closed_at,omitempty"`
	CreatedAt   *time.Time                 `json:"created_at,omitempty"`
	CustomerID  *string                    `json:"customer_id,omitempty"`
	Description *string                    `json:"description,omitempty"`
	ID          *string                    `json:"id,omitempty"`
	Priority    *string                    `json:"priority,omitempty"`
	Raw         PropertyTicketingTicketRaw `json:"raw"`
	Source      *string                    `json:"source,omitempty"`
	SourceRef   *string                    `json:"source_ref,omitempty"`
	Status      *TicketingTicketStatus     `json:"status,omitempty"`
	Subject     *string                    `json:"subject,omitempty"`
	Tags        []string                   `json:"tags,omitempty"`
	UpdatedAt   *time.Time                 `json:"updated_at,omitempty"`
}

func (t TicketingTicket) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TicketingTicket) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TicketingTicket) GetCategory() *string {
	if o == nil {
		return nil
	}
	return o.Category
}

func (o *TicketingTicket) GetClosedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ClosedAt
}

func (o *TicketingTicket) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *TicketingTicket) GetCustomerID() *string {
	if o == nil {
		return nil
	}
	return o.CustomerID
}

func (o *TicketingTicket) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *TicketingTicket) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *TicketingTicket) GetPriority() *string {
	if o == nil {
		return nil
	}
	return o.Priority
}

func (o *TicketingTicket) GetRaw() PropertyTicketingTicketRaw {
	if o == nil {
		return PropertyTicketingTicketRaw{}
	}
	return o.Raw
}

func (o *TicketingTicket) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *TicketingTicket) GetSourceRef() *string {
	if o == nil {
		return nil
	}
	return o.SourceRef
}

func (o *TicketingTicket) GetStatus() *TicketingTicketStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *TicketingTicket) GetSubject() *string {
	if o == nil {
		return nil
	}
	return o.Subject
}

func (o *TicketingTicket) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *TicketingTicket) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

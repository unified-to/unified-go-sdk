// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type TicketingEmailType string

const (
	TicketingEmailTypeWork  TicketingEmailType = "WORK"
	TicketingEmailTypeHome  TicketingEmailType = "HOME"
	TicketingEmailTypeOther TicketingEmailType = "OTHER"
)

func (e TicketingEmailType) ToPointer() *TicketingEmailType {
	return &e
}
func (e *TicketingEmailType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "WORK":
		fallthrough
	case "HOME":
		fallthrough
	case "OTHER":
		*e = TicketingEmailType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TicketingEmailType: %v", v)
	}
}

type TicketingEmail struct {
	Email string              `json:"email"`
	Type  *TicketingEmailType `json:"type,omitempty"`
}

func (o *TicketingEmail) GetEmail() string {
	if o == nil {
		return ""
	}
	return o.Email
}

func (o *TicketingEmail) GetType() *TicketingEmailType {
	if o == nil {
		return nil
	}
	return o.Type
}

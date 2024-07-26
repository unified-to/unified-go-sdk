// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type TicketingTelephoneType string

const (
	TicketingTelephoneTypeWork   TicketingTelephoneType = "WORK"
	TicketingTelephoneTypeHome   TicketingTelephoneType = "HOME"
	TicketingTelephoneTypeOther  TicketingTelephoneType = "OTHER"
	TicketingTelephoneTypeFax    TicketingTelephoneType = "FAX"
	TicketingTelephoneTypeMobile TicketingTelephoneType = "MOBILE"
)

func (e TicketingTelephoneType) ToPointer() *TicketingTelephoneType {
	return &e
}
func (e *TicketingTelephoneType) UnmarshalJSON(data []byte) error {
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
		fallthrough
	case "FAX":
		fallthrough
	case "MOBILE":
		*e = TicketingTelephoneType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TicketingTelephoneType: %v", v)
	}
}

type TicketingTelephone struct {
	Telephone string                  `json:"telephone"`
	Type      *TicketingTelephoneType `json:"type,omitempty"`
}

func (o *TicketingTelephone) GetTelephone() string {
	if o == nil {
		return ""
	}
	return o.Telephone
}

func (o *TicketingTelephone) GetType() *TicketingTelephoneType {
	if o == nil {
		return nil
	}
	return o.Type
}

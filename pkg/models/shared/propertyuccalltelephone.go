// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PropertyUcCallTelephoneType string

const (
	PropertyUcCallTelephoneTypeWork   PropertyUcCallTelephoneType = "WORK"
	PropertyUcCallTelephoneTypeHome   PropertyUcCallTelephoneType = "HOME"
	PropertyUcCallTelephoneTypeOther  PropertyUcCallTelephoneType = "OTHER"
	PropertyUcCallTelephoneTypeFax    PropertyUcCallTelephoneType = "FAX"
	PropertyUcCallTelephoneTypeMobile PropertyUcCallTelephoneType = "MOBILE"
)

func (e PropertyUcCallTelephoneType) ToPointer() *PropertyUcCallTelephoneType {
	return &e
}
func (e *PropertyUcCallTelephoneType) UnmarshalJSON(data []byte) error {
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
		*e = PropertyUcCallTelephoneType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PropertyUcCallTelephoneType: %v", v)
	}
}

// PropertyUcCallTelephone - The telephone number called
type PropertyUcCallTelephone struct {
	Telephone string                       `json:"telephone"`
	Type      *PropertyUcCallTelephoneType `json:"type,omitempty"`
}

func (o *PropertyUcCallTelephone) GetTelephone() string {
	if o == nil {
		return ""
	}
	return o.Telephone
}

func (o *PropertyUcCallTelephone) GetType() *PropertyUcCallTelephoneType {
	if o == nil {
		return nil
	}
	return o.Type
}

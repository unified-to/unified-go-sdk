// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type EnrichTelephoneType string

const (
	EnrichTelephoneTypeWork   EnrichTelephoneType = "WORK"
	EnrichTelephoneTypeHome   EnrichTelephoneType = "HOME"
	EnrichTelephoneTypeOther  EnrichTelephoneType = "OTHER"
	EnrichTelephoneTypeFax    EnrichTelephoneType = "FAX"
	EnrichTelephoneTypeMobile EnrichTelephoneType = "MOBILE"
)

func (e EnrichTelephoneType) ToPointer() *EnrichTelephoneType {
	return &e
}
func (e *EnrichTelephoneType) UnmarshalJSON(data []byte) error {
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
		*e = EnrichTelephoneType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EnrichTelephoneType: %v", v)
	}
}

type EnrichTelephone struct {
	Telephone string               `json:"telephone"`
	Type      *EnrichTelephoneType `json:"type,omitempty"`
}

func (o *EnrichTelephone) GetTelephone() string {
	if o == nil {
		return ""
	}
	return o.Telephone
}

func (o *EnrichTelephone) GetType() *EnrichTelephoneType {
	if o == nil {
		return nil
	}
	return o.Type
}

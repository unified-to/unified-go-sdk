// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type UcTelephoneType string

const (
	UcTelephoneTypeWork   UcTelephoneType = "WORK"
	UcTelephoneTypeHome   UcTelephoneType = "HOME"
	UcTelephoneTypeOther  UcTelephoneType = "OTHER"
	UcTelephoneTypeFax    UcTelephoneType = "FAX"
	UcTelephoneTypeMobile UcTelephoneType = "MOBILE"
)

func (e UcTelephoneType) ToPointer() *UcTelephoneType {
	return &e
}

func (e *UcTelephoneType) UnmarshalJSON(data []byte) error {
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
		*e = UcTelephoneType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UcTelephoneType: %v", v)
	}
}

type UcTelephone struct {
	Telephone string           `json:"telephone"`
	Type      *UcTelephoneType `json:"type,omitempty"`
}

func (o *UcTelephone) GetTelephone() string {
	if o == nil {
		return ""
	}
	return o.Telephone
}

func (o *UcTelephone) GetType() *UcTelephoneType {
	if o == nil {
		return nil
	}
	return o.Type
}

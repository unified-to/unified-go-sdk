// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type CrmTelephoneType string

const (
	CrmTelephoneTypeWork   CrmTelephoneType = "WORK"
	CrmTelephoneTypeHome   CrmTelephoneType = "HOME"
	CrmTelephoneTypeOther  CrmTelephoneType = "OTHER"
	CrmTelephoneTypeFax    CrmTelephoneType = "FAX"
	CrmTelephoneTypeMobile CrmTelephoneType = "MOBILE"
)

func (e CrmTelephoneType) ToPointer() *CrmTelephoneType {
	return &e
}
func (e *CrmTelephoneType) UnmarshalJSON(data []byte) error {
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
		*e = CrmTelephoneType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CrmTelephoneType: %v", v)
	}
}

type CrmTelephone struct {
	Telephone string            `json:"telephone"`
	Type      *CrmTelephoneType `json:"type,omitempty"`
}

func (o *CrmTelephone) GetTelephone() string {
	if o == nil {
		return ""
	}
	return o.Telephone
}

func (o *CrmTelephone) GetType() *CrmTelephoneType {
	if o == nil {
		return nil
	}
	return o.Type
}

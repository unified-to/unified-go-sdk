// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type AccountingTelephoneType string

const (
	AccountingTelephoneTypeWork   AccountingTelephoneType = "WORK"
	AccountingTelephoneTypeHome   AccountingTelephoneType = "HOME"
	AccountingTelephoneTypeOther  AccountingTelephoneType = "OTHER"
	AccountingTelephoneTypeFax    AccountingTelephoneType = "FAX"
	AccountingTelephoneTypeMobile AccountingTelephoneType = "MOBILE"
)

func (e AccountingTelephoneType) ToPointer() *AccountingTelephoneType {
	return &e
}
func (e *AccountingTelephoneType) UnmarshalJSON(data []byte) error {
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
		*e = AccountingTelephoneType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountingTelephoneType: %v", v)
	}
}

type AccountingTelephone struct {
	Telephone *string                  `json:"telephone,omitempty"`
	Type      *AccountingTelephoneType `json:"type,omitempty"`
}

func (o *AccountingTelephone) GetTelephone() *string {
	if o == nil {
		return nil
	}
	return o.Telephone
}

func (o *AccountingTelephone) GetType() *AccountingTelephoneType {
	if o == nil {
		return nil
	}
	return o.Type
}

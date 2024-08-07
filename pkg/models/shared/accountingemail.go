// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type AccountingEmailType string

const (
	AccountingEmailTypeWork  AccountingEmailType = "WORK"
	AccountingEmailTypeHome  AccountingEmailType = "HOME"
	AccountingEmailTypeOther AccountingEmailType = "OTHER"
)

func (e AccountingEmailType) ToPointer() *AccountingEmailType {
	return &e
}
func (e *AccountingEmailType) UnmarshalJSON(data []byte) error {
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
		*e = AccountingEmailType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountingEmailType: %v", v)
	}
}

type AccountingEmail struct {
	Email *string              `json:"email,omitempty"`
	Type  *AccountingEmailType `json:"type,omitempty"`
}

func (o *AccountingEmail) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *AccountingEmail) GetType() *AccountingEmailType {
	if o == nil {
		return nil
	}
	return o.Type
}

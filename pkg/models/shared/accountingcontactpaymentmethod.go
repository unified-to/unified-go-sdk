// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type AccountingContactPaymentMethodType string

const (
	AccountingContactPaymentMethodTypeAch     AccountingContactPaymentMethodType = "ACH"
	AccountingContactPaymentMethodTypeAlipay  AccountingContactPaymentMethodType = "ALIPAY"
	AccountingContactPaymentMethodTypeCard    AccountingContactPaymentMethodType = "CARD"
	AccountingContactPaymentMethodTypeGiropay AccountingContactPaymentMethodType = "GIROPAY"
	AccountingContactPaymentMethodTypeIdeal   AccountingContactPaymentMethodType = "IDEAL"
	AccountingContactPaymentMethodTypeOther   AccountingContactPaymentMethodType = "OTHER"
	AccountingContactPaymentMethodTypePaypal  AccountingContactPaymentMethodType = "PAYPAL"
)

func (e AccountingContactPaymentMethodType) ToPointer() *AccountingContactPaymentMethodType {
	return &e
}
func (e *AccountingContactPaymentMethodType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ACH":
		fallthrough
	case "ALIPAY":
		fallthrough
	case "CARD":
		fallthrough
	case "GIROPAY":
		fallthrough
	case "IDEAL":
		fallthrough
	case "OTHER":
		fallthrough
	case "PAYPAL":
		*e = AccountingContactPaymentMethodType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountingContactPaymentMethodType: %v", v)
	}
}

type AccountingContactPaymentMethod struct {
	ID   *string                            `json:"id,omitempty"`
	Name *string                            `json:"name,omitempty"`
	Type AccountingContactPaymentMethodType `json:"type"`
}

func (o *AccountingContactPaymentMethod) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingContactPaymentMethod) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AccountingContactPaymentMethod) GetType() AccountingContactPaymentMethodType {
	if o == nil {
		return AccountingContactPaymentMethodType("")
	}
	return o.Type
}

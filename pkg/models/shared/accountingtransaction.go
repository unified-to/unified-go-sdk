// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type AccountingTransactionType string

const (
	AccountingTransactionTypeReceive AccountingTransactionType = "RECEIVE"
	AccountingTransactionTypeSpend   AccountingTransactionType = "SPEND"
)

func (e AccountingTransactionType) ToPointer() *AccountingTransactionType {
	return &e
}

func (e *AccountingTransactionType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "RECEIVE":
		fallthrough
	case "SPEND":
		*e = AccountingTransactionType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountingTransactionType: %v", v)
	}
}

type AccountingTransaction struct {
	AccountID   string                            `json:"account_id"`
	CreatedAt   *string                           `json:"created_at,omitempty"`
	Currency    *string                           `json:"currency,omitempty"`
	Description *string                           `json:"description,omitempty"`
	ID          string                            `json:"id"`
	LineItems   []AccountingTransactionLineitem   `json:"line_items,omitempty"`
	Raw         *PropertyAccountingTransactionRaw `json:"raw,omitempty"`
	Reference   *string                           `json:"reference,omitempty"`
	TaxAmount   *float64                          `json:"tax_amount,omitempty"`
	TotalAmount float64                           `json:"total_amount"`
	Type        AccountingTransactionType         `json:"type"`
	UpdatedAt   *string                           `json:"updated_at,omitempty"`
}

func (o *AccountingTransaction) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *AccountingTransaction) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AccountingTransaction) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AccountingTransaction) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AccountingTransaction) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *AccountingTransaction) GetLineItems() []AccountingTransactionLineitem {
	if o == nil {
		return nil
	}
	return o.LineItems
}

func (o *AccountingTransaction) GetRaw() *PropertyAccountingTransactionRaw {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *AccountingTransaction) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *AccountingTransaction) GetTaxAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *AccountingTransaction) GetTotalAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.TotalAmount
}

func (o *AccountingTransaction) GetType() AccountingTransactionType {
	if o == nil {
		return AccountingTransactionType("")
	}
	return o.Type
}

func (o *AccountingTransaction) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
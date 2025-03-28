// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type AccountingTransactionRaw struct {
}

type AccountingTransaction struct {
	AccountID       *string                         `json:"account_id,omitempty"`
	ContactID       *string                         `json:"contact_id,omitempty"`
	Contacts        []AccountingTransactionContact  `json:"contacts,omitempty"`
	CreatedAt       *time.Time                      `json:"created_at,omitempty"`
	Currency        *string                         `json:"currency,omitempty"`
	CustomerMessage *string                         `json:"customer_message,omitempty"`
	ID              *string                         `json:"id,omitempty"`
	Lineitems       []AccountingTransactionLineItem `json:"lineitems,omitempty"`
	Memo            *string                         `json:"memo,omitempty"`
	PaymentMethod   *string                         `json:"payment_method,omitempty"`
	PaymentTerms    *string                         `json:"payment_terms,omitempty"`
	Raw             *AccountingTransactionRaw       `json:"raw,omitempty"`
	Reference       *string                         `json:"reference,omitempty"`
	SplitAccountID  *string                         `json:"split_account_id,omitempty"`
	SubTotalAmount  *float64                        `json:"sub_total_amount,omitempty"`
	TaxAmount       *float64                        `json:"tax_amount,omitempty"`
	TotalAmount     *float64                        `json:"total_amount,omitempty"`
	Type            *string                         `json:"type,omitempty"`
	UpdatedAt       *time.Time                      `json:"updated_at,omitempty"`
}

func (a AccountingTransaction) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingTransaction) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingTransaction) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *AccountingTransaction) GetContactID() *string {
	if o == nil {
		return nil
	}
	return o.ContactID
}

func (o *AccountingTransaction) GetContacts() []AccountingTransactionContact {
	if o == nil {
		return nil
	}
	return o.Contacts
}

func (o *AccountingTransaction) GetCreatedAt() *time.Time {
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

func (o *AccountingTransaction) GetCustomerMessage() *string {
	if o == nil {
		return nil
	}
	return o.CustomerMessage
}

func (o *AccountingTransaction) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingTransaction) GetLineitems() []AccountingTransactionLineItem {
	if o == nil {
		return nil
	}
	return o.Lineitems
}

func (o *AccountingTransaction) GetMemo() *string {
	if o == nil {
		return nil
	}
	return o.Memo
}

func (o *AccountingTransaction) GetPaymentMethod() *string {
	if o == nil {
		return nil
	}
	return o.PaymentMethod
}

func (o *AccountingTransaction) GetPaymentTerms() *string {
	if o == nil {
		return nil
	}
	return o.PaymentTerms
}

func (o *AccountingTransaction) GetRaw() *AccountingTransactionRaw {
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

func (o *AccountingTransaction) GetSplitAccountID() *string {
	if o == nil {
		return nil
	}
	return o.SplitAccountID
}

func (o *AccountingTransaction) GetSubTotalAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.SubTotalAmount
}

func (o *AccountingTransaction) GetTaxAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *AccountingTransaction) GetTotalAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *AccountingTransaction) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *AccountingTransaction) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type AccountingPayment struct {
	AccountID     *string                       `json:"account_id,omitempty"`
	ContactID     *string                       `json:"contact_id,omitempty"`
	CreatedAt     *time.Time                    `json:"created_at,omitempty"`
	Currency      *string                       `default:"USD" json:"currency"`
	ID            *string                       `json:"id,omitempty"`
	InvoiceID     *string                       `json:"invoice_id,omitempty"`
	Notes         *string                       `json:"notes,omitempty"`
	PaymentMethod *string                       `json:"payment_method,omitempty"`
	Raw           *PropertyAccountingPaymentRaw `json:"raw,omitempty"`
	Reference     *string                       `json:"reference,omitempty"`
	TotalAmount   *float64                      `json:"total_amount,omitempty"`
	UpdatedAt     *time.Time                    `json:"updated_at,omitempty"`
}

func (a AccountingPayment) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingPayment) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingPayment) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *AccountingPayment) GetContactID() *string {
	if o == nil {
		return nil
	}
	return o.ContactID
}

func (o *AccountingPayment) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AccountingPayment) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AccountingPayment) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingPayment) GetInvoiceID() *string {
	if o == nil {
		return nil
	}
	return o.InvoiceID
}

func (o *AccountingPayment) GetNotes() *string {
	if o == nil {
		return nil
	}
	return o.Notes
}

func (o *AccountingPayment) GetPaymentMethod() *string {
	if o == nil {
		return nil
	}
	return o.PaymentMethod
}

func (o *AccountingPayment) GetRaw() *PropertyAccountingPaymentRaw {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *AccountingPayment) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *AccountingPayment) GetTotalAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *AccountingPayment) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

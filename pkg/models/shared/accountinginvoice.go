// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type PaymentCollectionMethod string

const (
	PaymentCollectionMethodSendInvoice         PaymentCollectionMethod = "send_invoice"
	PaymentCollectionMethodChargeAutomatically PaymentCollectionMethod = "charge_automatically"
)

func (e PaymentCollectionMethod) ToPointer() *PaymentCollectionMethod {
	return &e
}
func (e *PaymentCollectionMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "send_invoice":
		fallthrough
	case "charge_automatically":
		*e = PaymentCollectionMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PaymentCollectionMethod: %v", v)
	}
}

type AccountingInvoiceStatus string

const (
	AccountingInvoiceStatusDraft             AccountingInvoiceStatus = "DRAFT"
	AccountingInvoiceStatusVoided            AccountingInvoiceStatus = "VOIDED"
	AccountingInvoiceStatusAuthorized        AccountingInvoiceStatus = "AUTHORIZED"
	AccountingInvoiceStatusPaid              AccountingInvoiceStatus = "PAID"
	AccountingInvoiceStatusPartiallyPaid     AccountingInvoiceStatus = "PARTIALLY_PAID"
	AccountingInvoiceStatusPartiallyRefunded AccountingInvoiceStatus = "PARTIALLY_REFUNDED"
	AccountingInvoiceStatusRefunded          AccountingInvoiceStatus = "REFUNDED"
)

func (e AccountingInvoiceStatus) ToPointer() *AccountingInvoiceStatus {
	return &e
}
func (e *AccountingInvoiceStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "DRAFT":
		fallthrough
	case "VOIDED":
		fallthrough
	case "AUTHORIZED":
		fallthrough
	case "PAID":
		fallthrough
	case "PARTIALLY_PAID":
		fallthrough
	case "PARTIALLY_REFUNDED":
		fallthrough
	case "REFUNDED":
		*e = AccountingInvoiceStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountingInvoiceStatus: %v", v)
	}
}

type AccountingInvoiceType string

const (
	AccountingInvoiceTypeBill       AccountingInvoiceType = "BILL"
	AccountingInvoiceTypeInvoice    AccountingInvoiceType = "INVOICE"
	AccountingInvoiceTypeCreditmemo AccountingInvoiceType = "CREDITMEMO"
)

func (e AccountingInvoiceType) ToPointer() *AccountingInvoiceType {
	return &e
}
func (e *AccountingInvoiceType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "BILL":
		fallthrough
	case "INVOICE":
		fallthrough
	case "CREDITMEMO":
		*e = AccountingInvoiceType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AccountingInvoiceType: %v", v)
	}
}

type AccountingInvoice struct {
	Attachments             []AccountingAttachment   `json:"attachments,omitempty"`
	BalanceAmount           *float64                 `json:"balance_amount,omitempty"`
	CancelledAt             *time.Time               `json:"cancelled_at,omitempty"`
	ContactID               *string                  `json:"contact_id,omitempty"`
	CreatedAt               *time.Time               `json:"created_at,omitempty"`
	Currency                *string                  `json:"currency,omitempty"`
	DiscountAmount          *float64                 `json:"discount_amount,omitempty"`
	DueAt                   *time.Time               `json:"due_at,omitempty"`
	ID                      *string                  `json:"id,omitempty"`
	InvoiceAt               *time.Time               `json:"invoice_at,omitempty"`
	InvoiceNumber           *string                  `json:"invoice_number,omitempty"`
	Lineitems               []AccountingLineitem     `json:"lineitems,omitempty"`
	Notes                   *string                  `json:"notes,omitempty"`
	PaidAmount              *float64                 `json:"paid_amount,omitempty"`
	PaidAt                  *time.Time               `json:"paid_at,omitempty"`
	PaymentCollectionMethod *PaymentCollectionMethod `json:"payment_collection_method,omitempty"`
	PostedAt                *time.Time               `json:"posted_at,omitempty"`
	Raw                     map[string]any           `json:"raw,omitempty"`
	RefundAmount            *float64                 `json:"refund_amount,omitempty"`
	RefundReason            *string                  `json:"refund_reason,omitempty"`
	RefundedAt              *time.Time               `json:"refunded_at,omitempty"`
	Send                    *bool                    `json:"send,omitempty"`
	Status                  *AccountingInvoiceStatus `json:"status,omitempty"`
	TaxAmount               *float64                 `json:"tax_amount,omitempty"`
	TotalAmount             *float64                 `json:"total_amount,omitempty"`
	Type                    *AccountingInvoiceType   `json:"type,omitempty"`
	UpdatedAt               *time.Time               `json:"updated_at,omitempty"`
	URL                     *string                  `json:"url,omitempty"`
}

func (a AccountingInvoice) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingInvoice) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingInvoice) GetAttachments() []AccountingAttachment {
	if o == nil {
		return nil
	}
	return o.Attachments
}

func (o *AccountingInvoice) GetBalanceAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.BalanceAmount
}

func (o *AccountingInvoice) GetCancelledAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CancelledAt
}

func (o *AccountingInvoice) GetContactID() *string {
	if o == nil {
		return nil
	}
	return o.ContactID
}

func (o *AccountingInvoice) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AccountingInvoice) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AccountingInvoice) GetDiscountAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.DiscountAmount
}

func (o *AccountingInvoice) GetDueAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.DueAt
}

func (o *AccountingInvoice) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingInvoice) GetInvoiceAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.InvoiceAt
}

func (o *AccountingInvoice) GetInvoiceNumber() *string {
	if o == nil {
		return nil
	}
	return o.InvoiceNumber
}

func (o *AccountingInvoice) GetLineitems() []AccountingLineitem {
	if o == nil {
		return nil
	}
	return o.Lineitems
}

func (o *AccountingInvoice) GetNotes() *string {
	if o == nil {
		return nil
	}
	return o.Notes
}

func (o *AccountingInvoice) GetPaidAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.PaidAmount
}

func (o *AccountingInvoice) GetPaidAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.PaidAt
}

func (o *AccountingInvoice) GetPaymentCollectionMethod() *PaymentCollectionMethod {
	if o == nil {
		return nil
	}
	return o.PaymentCollectionMethod
}

func (o *AccountingInvoice) GetPostedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.PostedAt
}

func (o *AccountingInvoice) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *AccountingInvoice) GetRefundAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.RefundAmount
}

func (o *AccountingInvoice) GetRefundReason() *string {
	if o == nil {
		return nil
	}
	return o.RefundReason
}

func (o *AccountingInvoice) GetRefundedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RefundedAt
}

func (o *AccountingInvoice) GetSend() *bool {
	if o == nil {
		return nil
	}
	return o.Send
}

func (o *AccountingInvoice) GetStatus() *AccountingInvoiceStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *AccountingInvoice) GetTaxAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *AccountingInvoice) GetTotalAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *AccountingInvoice) GetType() *AccountingInvoiceType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *AccountingInvoice) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *AccountingInvoice) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

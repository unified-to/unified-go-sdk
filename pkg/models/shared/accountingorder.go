// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type AccountingOrderStatus string

const (
	AccountingOrderStatusDraft             AccountingOrderStatus = "DRAFT"
	AccountingOrderStatusVoided            AccountingOrderStatus = "VOIDED"
	AccountingOrderStatusAuthorized        AccountingOrderStatus = "AUTHORIZED"
	AccountingOrderStatusPaid              AccountingOrderStatus = "PAID"
	AccountingOrderStatusPartiallyPaid     AccountingOrderStatus = "PARTIALLY_PAID"
	AccountingOrderStatusPartiallyRefunded AccountingOrderStatus = "PARTIALLY_REFUNDED"
	AccountingOrderStatusRefunded          AccountingOrderStatus = "REFUNDED"
)

func (e AccountingOrderStatus) ToPointer() *AccountingOrderStatus {
	return &e
}

type AccountingOrderType string

const (
	AccountingOrderTypeSales    AccountingOrderType = "SALES"
	AccountingOrderTypePurchase AccountingOrderType = "PURCHASE"
)

func (e AccountingOrderType) ToPointer() *AccountingOrderType {
	return &e
}

type AccountingOrder struct {
	AccountID       *string                                 `json:"account_id,omitempty"`
	BillingAddress  *PropertyAccountingOrderBillingAddress  `json:"billing_address,omitempty"`
	ContactID       *string                                 `json:"contact_id,omitempty"`
	CreatedAt       *time.Time                              `json:"created_at,omitempty"`
	Currency        *string                                 `json:"currency,omitempty"`
	ID              *string                                 `json:"id,omitempty"`
	Lineitems       []AccountingLineitem                    `json:"lineitems,omitempty"`
	PostedAt        *time.Time                              `json:"posted_at,omitempty"`
	Raw             map[string]any                          `json:"raw,omitempty"`
	ShippingAddress *PropertyAccountingOrderShippingAddress `json:"shipping_address,omitempty"`
	Status          *AccountingOrderStatus                  `json:"status,omitempty"`
	TotalAmount     *float64                                `json:"total_amount,omitempty"`
	Type            *AccountingOrderType                    `json:"type,omitempty"`
	UpdatedAt       *time.Time                              `json:"updated_at,omitempty"`
}

func (a AccountingOrder) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingOrder) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingOrder) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *AccountingOrder) GetBillingAddress() *PropertyAccountingOrderBillingAddress {
	if o == nil {
		return nil
	}
	return o.BillingAddress
}

func (o *AccountingOrder) GetContactID() *string {
	if o == nil {
		return nil
	}
	return o.ContactID
}

func (o *AccountingOrder) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AccountingOrder) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AccountingOrder) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingOrder) GetLineitems() []AccountingLineitem {
	if o == nil {
		return nil
	}
	return o.Lineitems
}

func (o *AccountingOrder) GetPostedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.PostedAt
}

func (o *AccountingOrder) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *AccountingOrder) GetShippingAddress() *PropertyAccountingOrderShippingAddress {
	if o == nil {
		return nil
	}
	return o.ShippingAddress
}

func (o *AccountingOrder) GetStatus() *AccountingOrderStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *AccountingOrder) GetTotalAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *AccountingOrder) GetType() *AccountingOrderType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *AccountingOrder) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

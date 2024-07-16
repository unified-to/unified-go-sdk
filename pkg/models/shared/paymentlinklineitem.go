// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type PaymentLinkLineitem struct {
	AccountID       *string    `json:"account_id,omitempty"`
	CreatedAt       *time.Time `json:"created_at,omitempty"`
	DiscountAmount  *float64   `json:"discount_amount,omitempty"`
	ID              *string    `json:"id,omitempty"`
	ItemDescription *string    `json:"item_description,omitempty"`
	ItemID          *string    `json:"item_id,omitempty"`
	ItemName        *string    `json:"item_name,omitempty"`
	ItemSku         *string    `json:"item_sku,omitempty"`
	Notes           *string    `json:"notes,omitempty"`
	RefundAmount    *float64   `json:"refund_amount,omitempty"`
	RefundedAt      *time.Time `json:"refunded_at,omitempty"`
	TaxAmount       *float64   `json:"tax_amount,omitempty"`
	TaxrateID       *string    `json:"taxrate_id,omitempty"`
	TotalAmount     *float64   `json:"total_amount,omitempty"`
	UnitAmount      *float64   `json:"unit_amount,omitempty"`
	UnitQuantity    *float64   `json:"unit_quantity,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
}

func (p PaymentLinkLineitem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PaymentLinkLineitem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PaymentLinkLineitem) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *PaymentLinkLineitem) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *PaymentLinkLineitem) GetDiscountAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.DiscountAmount
}

func (o *PaymentLinkLineitem) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PaymentLinkLineitem) GetItemDescription() *string {
	if o == nil {
		return nil
	}
	return o.ItemDescription
}

func (o *PaymentLinkLineitem) GetItemID() *string {
	if o == nil {
		return nil
	}
	return o.ItemID
}

func (o *PaymentLinkLineitem) GetItemName() *string {
	if o == nil {
		return nil
	}
	return o.ItemName
}

func (o *PaymentLinkLineitem) GetItemSku() *string {
	if o == nil {
		return nil
	}
	return o.ItemSku
}

func (o *PaymentLinkLineitem) GetNotes() *string {
	if o == nil {
		return nil
	}
	return o.Notes
}

func (o *PaymentLinkLineitem) GetRefundAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.RefundAmount
}

func (o *PaymentLinkLineitem) GetRefundedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RefundedAt
}

func (o *PaymentLinkLineitem) GetTaxAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *PaymentLinkLineitem) GetTaxrateID() *string {
	if o == nil {
		return nil
	}
	return o.TaxrateID
}

func (o *PaymentLinkLineitem) GetTotalAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalAmount
}

func (o *PaymentLinkLineitem) GetUnitAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.UnitAmount
}

func (o *PaymentLinkLineitem) GetUnitQuantity() *float64 {
	if o == nil {
		return nil
	}
	return o.UnitQuantity
}

func (o *PaymentLinkLineitem) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

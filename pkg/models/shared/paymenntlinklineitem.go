// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type PaymenntLinkLineitem struct {
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
	TotalAmount     float64    `json:"total_amount"`
	UnitAmount      *float64   `json:"unit_amount,omitempty"`
	UnitQuantity    *float64   `json:"unit_quantity,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at,omitempty"`
}

func (p PaymenntLinkLineitem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PaymenntLinkLineitem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PaymenntLinkLineitem) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *PaymenntLinkLineitem) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *PaymenntLinkLineitem) GetDiscountAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.DiscountAmount
}

func (o *PaymenntLinkLineitem) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PaymenntLinkLineitem) GetItemDescription() *string {
	if o == nil {
		return nil
	}
	return o.ItemDescription
}

func (o *PaymenntLinkLineitem) GetItemID() *string {
	if o == nil {
		return nil
	}
	return o.ItemID
}

func (o *PaymenntLinkLineitem) GetItemName() *string {
	if o == nil {
		return nil
	}
	return o.ItemName
}

func (o *PaymenntLinkLineitem) GetItemSku() *string {
	if o == nil {
		return nil
	}
	return o.ItemSku
}

func (o *PaymenntLinkLineitem) GetNotes() *string {
	if o == nil {
		return nil
	}
	return o.Notes
}

func (o *PaymenntLinkLineitem) GetRefundAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.RefundAmount
}

func (o *PaymenntLinkLineitem) GetRefundedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.RefundedAt
}

func (o *PaymenntLinkLineitem) GetTaxAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TaxAmount
}

func (o *PaymenntLinkLineitem) GetTaxrateID() *string {
	if o == nil {
		return nil
	}
	return o.TaxrateID
}

func (o *PaymenntLinkLineitem) GetTotalAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.TotalAmount
}

func (o *PaymenntLinkLineitem) GetUnitAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.UnitAmount
}

func (o *PaymenntLinkLineitem) GetUnitQuantity() *float64 {
	if o == nil {
		return nil
	}
	return o.UnitQuantity
}

func (o *PaymenntLinkLineitem) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
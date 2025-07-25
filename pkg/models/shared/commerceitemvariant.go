// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type SizeUnit string

const (
	SizeUnitCm   SizeUnit = "cm"
	SizeUnitInch SizeUnit = "inch"
)

func (e SizeUnit) ToPointer() *SizeUnit {
	return &e
}

type WeightUnit string

const (
	WeightUnitG  WeightUnit = "g"
	WeightUnitKg WeightUnit = "kg"
	WeightUnitOz WeightUnit = "oz"
	WeightUnitLb WeightUnit = "lb"
)

func (e WeightUnit) ToPointer() *WeightUnit {
	return &e
}

type CommerceItemVariant struct {
	AvailableAt       *time.Time           `json:"available_at,omitempty"`
	Description       *string              `json:"description,omitempty"`
	Height            *float64             `json:"height,omitempty"`
	ID                *string              `json:"id,omitempty"`
	InventoryID       *string              `json:"inventory_id,omitempty"`
	IsActive          *bool                `json:"is_active,omitempty"`
	IsFeatured        *bool                `json:"is_featured,omitempty"`
	IsVisible         *bool                `json:"is_visible,omitempty"`
	Length            *float64             `json:"length,omitempty"`
	Media             []CommerceItemMedia  `json:"media,omitempty"`
	Metadata          []CommerceMetadata   `json:"metadata,omitempty"`
	Name              *string              `json:"name,omitempty"`
	Options           []CommerceItemOption `json:"options,omitempty"`
	Prices            []CommerceItemPrice  `json:"prices,omitempty"`
	PublicDescription *string              `json:"public_description,omitempty"`
	PublicName        *string              `json:"public_name,omitempty"`
	RequiresShipping  *bool                `json:"requires_shipping,omitempty"`
	SizeUnit          *SizeUnit            `json:"size_unit,omitempty"`
	Sku               *string              `json:"sku,omitempty"`
	Tags              []string             `json:"tags,omitempty"`
	TotalStock        *float64             `json:"total_stock,omitempty"`
	Weight            *float64             `json:"weight,omitempty"`
	WeightUnit        *WeightUnit          `json:"weight_unit,omitempty"`
	Width             *float64             `json:"width,omitempty"`
}

func (c CommerceItemVariant) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CommerceItemVariant) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CommerceItemVariant) GetAvailableAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.AvailableAt
}

func (o *CommerceItemVariant) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CommerceItemVariant) GetHeight() *float64 {
	if o == nil {
		return nil
	}
	return o.Height
}

func (o *CommerceItemVariant) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CommerceItemVariant) GetInventoryID() *string {
	if o == nil {
		return nil
	}
	return o.InventoryID
}

func (o *CommerceItemVariant) GetIsActive() *bool {
	if o == nil {
		return nil
	}
	return o.IsActive
}

func (o *CommerceItemVariant) GetIsFeatured() *bool {
	if o == nil {
		return nil
	}
	return o.IsFeatured
}

func (o *CommerceItemVariant) GetIsVisible() *bool {
	if o == nil {
		return nil
	}
	return o.IsVisible
}

func (o *CommerceItemVariant) GetLength() *float64 {
	if o == nil {
		return nil
	}
	return o.Length
}

func (o *CommerceItemVariant) GetMedia() []CommerceItemMedia {
	if o == nil {
		return nil
	}
	return o.Media
}

func (o *CommerceItemVariant) GetMetadata() []CommerceMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *CommerceItemVariant) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CommerceItemVariant) GetOptions() []CommerceItemOption {
	if o == nil {
		return nil
	}
	return o.Options
}

func (o *CommerceItemVariant) GetPrices() []CommerceItemPrice {
	if o == nil {
		return nil
	}
	return o.Prices
}

func (o *CommerceItemVariant) GetPublicDescription() *string {
	if o == nil {
		return nil
	}
	return o.PublicDescription
}

func (o *CommerceItemVariant) GetPublicName() *string {
	if o == nil {
		return nil
	}
	return o.PublicName
}

func (o *CommerceItemVariant) GetRequiresShipping() *bool {
	if o == nil {
		return nil
	}
	return o.RequiresShipping
}

func (o *CommerceItemVariant) GetSizeUnit() *SizeUnit {
	if o == nil {
		return nil
	}
	return o.SizeUnit
}

func (o *CommerceItemVariant) GetSku() *string {
	if o == nil {
		return nil
	}
	return o.Sku
}

func (o *CommerceItemVariant) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CommerceItemVariant) GetTotalStock() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalStock
}

func (o *CommerceItemVariant) GetWeight() *float64 {
	if o == nil {
		return nil
	}
	return o.Weight
}

func (o *CommerceItemVariant) GetWeightUnit() *WeightUnit {
	if o == nil {
		return nil
	}
	return o.WeightUnit
}

func (o *CommerceItemVariant) GetWidth() *float64 {
	if o == nil {
		return nil
	}
	return o.Width
}

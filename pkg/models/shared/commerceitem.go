// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type CommerceItem struct {
	CreatedAt         *time.Time          `json:"created_at,omitempty"`
	Description       *string             `json:"description,omitempty"`
	ID                *string             `json:"id,omitempty"`
	IsActive          *bool               `json:"is_active,omitempty"`
	IsTaxable         *bool               `json:"is_taxable,omitempty"`
	Media             []CommerceItemMedia `json:"media,omitempty"`
	Name              *string             `json:"name,omitempty"`
	PublicDescription *string             `json:"public_description,omitempty"`
	PublicName        *string             `json:"public_name,omitempty"`
	Raw               map[string]any      `json:"raw,omitempty"`
	Slug              *string             `json:"slug,omitempty"`
	Tags              []string            `json:"tags,omitempty"`
	Type              *string             `json:"type,omitempty"`
	UpdatedAt         *time.Time          `json:"updated_at,omitempty"`
	// first variant is the default variant
	Variants   []CommerceItemVariant `json:"variants,omitempty"`
	VendorName *string               `json:"vendor_name,omitempty"`
}

func (c CommerceItem) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CommerceItem) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CommerceItem) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CommerceItem) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CommerceItem) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CommerceItem) GetIsActive() *bool {
	if o == nil {
		return nil
	}
	return o.IsActive
}

func (o *CommerceItem) GetIsTaxable() *bool {
	if o == nil {
		return nil
	}
	return o.IsTaxable
}

func (o *CommerceItem) GetMedia() []CommerceItemMedia {
	if o == nil {
		return nil
	}
	return o.Media
}

func (o *CommerceItem) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CommerceItem) GetPublicDescription() *string {
	if o == nil {
		return nil
	}
	return o.PublicDescription
}

func (o *CommerceItem) GetPublicName() *string {
	if o == nil {
		return nil
	}
	return o.PublicName
}

func (o *CommerceItem) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *CommerceItem) GetSlug() *string {
	if o == nil {
		return nil
	}
	return o.Slug
}

func (o *CommerceItem) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CommerceItem) GetType() *string {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *CommerceItem) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *CommerceItem) GetVariants() []CommerceItemVariant {
	if o == nil {
		return nil
	}
	return o.Variants
}

func (o *CommerceItem) GetVendorName() *string {
	if o == nil {
		return nil
	}
	return o.VendorName
}

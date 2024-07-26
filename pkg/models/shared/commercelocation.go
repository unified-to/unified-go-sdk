// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type CommerceLocation struct {
	Address     *PropertyCommerceLocationAddress `json:"address,omitempty"`
	CreatedAt   *time.Time                       `json:"created_at,omitempty"`
	Description *string                          `json:"description,omitempty"`
	ID          *string                          `json:"id,omitempty"`
	IsActive    *bool                            `json:"is_active,omitempty"`
	Name        string                           `json:"name"`
	Raw         map[string]any                   `json:"raw,omitempty"`
	UpdatedAt   *time.Time                       `json:"updated_at,omitempty"`
}

func (c CommerceLocation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CommerceLocation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CommerceLocation) GetAddress() *PropertyCommerceLocationAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *CommerceLocation) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CommerceLocation) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CommerceLocation) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CommerceLocation) GetIsActive() *bool {
	if o == nil {
		return nil
	}
	return o.IsActive
}

func (o *CommerceLocation) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CommerceLocation) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *CommerceLocation) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

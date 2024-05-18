// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type AccountingTaxrate struct {
	CreatedAt   *time.Time     `json:"created_at,omitempty"`
	Description *string        `json:"description,omitempty"`
	ID          *string        `json:"id,omitempty"`
	IsActive    *bool          `json:"is_active,omitempty"`
	Name        *string        `json:"name,omitempty"`
	Rate        *float64       `json:"rate,omitempty"`
	Raw         map[string]any `json:"raw,omitempty"`
	UpdatedAt   *time.Time     `json:"updated_at,omitempty"`
}

func (a AccountingTaxrate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountingTaxrate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountingTaxrate) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AccountingTaxrate) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *AccountingTaxrate) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AccountingTaxrate) GetIsActive() *bool {
	if o == nil {
		return nil
	}
	return o.IsActive
}

func (o *AccountingTaxrate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AccountingTaxrate) GetRate() *float64 {
	if o == nil {
		return nil
	}
	return o.Rate
}

func (o *AccountingTaxrate) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *AccountingTaxrate) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

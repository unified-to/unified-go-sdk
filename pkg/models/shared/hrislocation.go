// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type HrisLocation struct {
	Address            *PropertyHrisLocationAddress `json:"address,omitempty"`
	CreatedAt          *time.Time                   `json:"created_at,omitempty"`
	Currency           *string                      `json:"currency,omitempty"`
	Description        *string                      `json:"description,omitempty"`
	ExternalIdentifier *string                      `json:"external_identifier,omitempty"`
	ID                 *string                      `json:"id,omitempty"`
	IsActive           *bool                        `json:"is_active,omitempty"`
	IsHq               *bool                        `json:"is_hq,omitempty"`
	LanguageLocale     *string                      `json:"language_locale,omitempty"`
	Name               *string                      `json:"name,omitempty"`
	ParentID           *string                      `json:"parent_id,omitempty"`
	Raw                map[string]any               `json:"raw,omitempty"`
	Telephones         []HrisTelephone              `json:"telephones,omitempty"`
	Timezone           *string                      `json:"timezone,omitempty"`
	UpdatedAt          *time.Time                   `json:"updated_at,omitempty"`
}

func (h HrisLocation) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HrisLocation) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *HrisLocation) GetAddress() *PropertyHrisLocationAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *HrisLocation) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *HrisLocation) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *HrisLocation) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *HrisLocation) GetExternalIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.ExternalIdentifier
}

func (o *HrisLocation) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *HrisLocation) GetIsActive() *bool {
	if o == nil {
		return nil
	}
	return o.IsActive
}

func (o *HrisLocation) GetIsHq() *bool {
	if o == nil {
		return nil
	}
	return o.IsHq
}

func (o *HrisLocation) GetLanguageLocale() *string {
	if o == nil {
		return nil
	}
	return o.LanguageLocale
}

func (o *HrisLocation) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *HrisLocation) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *HrisLocation) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *HrisLocation) GetTelephones() []HrisTelephone {
	if o == nil {
		return nil
	}
	return o.Telephones
}

func (o *HrisLocation) GetTimezone() *string {
	if o == nil {
		return nil
	}
	return o.Timezone
}

func (o *HrisLocation) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
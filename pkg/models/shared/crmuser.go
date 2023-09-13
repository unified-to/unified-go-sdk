// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

type CrmUser struct {
	Active         *bool                   `json:"active,omitempty"`
	Address        *PropertyCrmUserAddress `json:"address,omitempty"`
	CreatedAt      *types.Date             `json:"created_at,omitempty"`
	Currency       *string                 `json:"currency,omitempty"`
	Department     *string                 `json:"department,omitempty"`
	Division       *string                 `json:"division,omitempty"`
	Emails         []CrmEmail              `json:"emails,omitempty"`
	ID             *string                 `json:"id,omitempty"`
	ImageURL       *string                 `json:"image_url,omitempty"`
	LanguageLocale *string                 `json:"language_locale,omitempty"`
	Name           *string                 `json:"name,omitempty"`
	Raw            *PropertyCrmUserRaw     `json:"raw,omitempty"`
	Telephones     []CrmTelephone          `json:"telephones,omitempty"`
	Timezone       *string                 `json:"timezone,omitempty"`
	Title          *string                 `json:"title,omitempty"`
	UpdatedAt      *types.Date             `json:"updated_at,omitempty"`
}

func (o *CrmUser) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *CrmUser) GetAddress() *PropertyCrmUserAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *CrmUser) GetCreatedAt() *types.Date {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CrmUser) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *CrmUser) GetDepartment() *string {
	if o == nil {
		return nil
	}
	return o.Department
}

func (o *CrmUser) GetDivision() *string {
	if o == nil {
		return nil
	}
	return o.Division
}

func (o *CrmUser) GetEmails() []CrmEmail {
	if o == nil {
		return nil
	}
	return o.Emails
}

func (o *CrmUser) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CrmUser) GetImageURL() *string {
	if o == nil {
		return nil
	}
	return o.ImageURL
}

func (o *CrmUser) GetLanguageLocale() *string {
	if o == nil {
		return nil
	}
	return o.LanguageLocale
}

func (o *CrmUser) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CrmUser) GetRaw() *PropertyCrmUserRaw {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *CrmUser) GetTelephones() []CrmTelephone {
	if o == nil {
		return nil
	}
	return o.Telephones
}

func (o *CrmUser) GetTimezone() *string {
	if o == nil {
		return nil
	}
	return o.Timezone
}

func (o *CrmUser) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *CrmUser) GetUpdatedAt() *types.Date {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

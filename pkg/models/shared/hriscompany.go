// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type HrisCompany struct {
	Address   *PropertyHrisCompanyAddress `json:"address,omitempty"`
	CreatedAt *time.Time                  `json:"created_at,omitempty"`
	ID        *string                     `json:"id,omitempty"`
	LegalName *string                     `json:"legal_name,omitempty"`
	Name      *string                     `json:"name,omitempty"`
	Raw       map[string]any              `json:"raw,omitempty"`
	UpdatedAt *time.Time                  `json:"updated_at,omitempty"`
}

func (h HrisCompany) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HrisCompany) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *HrisCompany) GetAddress() *PropertyHrisCompanyAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *HrisCompany) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *HrisCompany) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *HrisCompany) GetLegalName() *string {
	if o == nil {
		return nil
	}
	return o.LegalName
}

func (o *HrisCompany) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *HrisCompany) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *HrisCompany) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

// CrmCompany - A company represents an organization that optionally is associated with a deal and/or contacts
type CrmCompany struct {
	Address   *PropertyCrmCompanyAddress `json:"address,omitempty"`
	CreatedAt *time.Time                 `json:"created_at,omitempty"`
	// An array of deal IDs associated with this contact
	DealIds  []string   `json:"deal_ids,omitempty"`
	Emails   []CrmEmail `json:"emails,omitempty"`
	ID       *string    `json:"id,omitempty"`
	IsActive *bool      `json:"is_active,omitempty"`
	Name     *string    `json:"name,omitempty"`
	// The raw data returned by the integration for this company
	Raw        *PropertyCrmCompanyRaw `json:"raw,omitempty"`
	Tags       []string               `json:"tags,omitempty"`
	Telephones []CrmTelephone         `json:"telephones,omitempty"`
	UpdatedAt  *time.Time             `json:"updated_at,omitempty"`
	Websites   []string               `json:"websites,omitempty"`
}

func (c CrmCompany) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CrmCompany) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CrmCompany) GetAddress() *PropertyCrmCompanyAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *CrmCompany) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CrmCompany) GetDealIds() []string {
	if o == nil {
		return nil
	}
	return o.DealIds
}

func (o *CrmCompany) GetEmails() []CrmEmail {
	if o == nil {
		return nil
	}
	return o.Emails
}

func (o *CrmCompany) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CrmCompany) GetIsActive() *bool {
	if o == nil {
		return nil
	}
	return o.IsActive
}

func (o *CrmCompany) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CrmCompany) GetRaw() *PropertyCrmCompanyRaw {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *CrmCompany) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CrmCompany) GetTelephones() []CrmTelephone {
	if o == nil {
		return nil
	}
	return o.Telephones
}

func (o *CrmCompany) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *CrmCompany) GetWebsites() []string {
	if o == nil {
		return nil
	}
	return o.Websites
}

// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type CrmContactRaw struct {
}

// CrmContact - A contact represents a person that optionally is associated with a deal and/or a company
type CrmContact struct {
	Address *PropertyCrmContactAddress `json:"address,omitempty"`
	Company *string                    `json:"company,omitempty"`
	// An array of company IDs associated with this contact
	CompanyIds []string   `json:"company_ids,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	// An array of deal IDs associated with this contact
	DealIds []string `json:"deal_ids,omitempty"`
	// An array of email addresses for this contact
	Emails []CrmEmail `json:"emails,omitempty"`
	ID     *string    `json:"id,omitempty"`
	// Additional URLs associated with the contact e.g., LinkedIn, website, etc
	LinkUrls []string       `json:"link_urls,omitempty"`
	Name     *string        `json:"name,omitempty"`
	Raw      *CrmContactRaw `json:"raw,omitempty"`
	// An array of telephones for this contact
	Telephones []CrmTelephone `json:"telephones,omitempty"`
	Title      *string        `json:"title,omitempty"`
	UpdatedAt  *time.Time     `json:"updated_at,omitempty"`
	UserID     *string        `json:"user_id,omitempty"`
}

func (c CrmContact) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CrmContact) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CrmContact) GetAddress() *PropertyCrmContactAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *CrmContact) GetCompany() *string {
	if o == nil {
		return nil
	}
	return o.Company
}

func (o *CrmContact) GetCompanyIds() []string {
	if o == nil {
		return nil
	}
	return o.CompanyIds
}

func (o *CrmContact) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CrmContact) GetDealIds() []string {
	if o == nil {
		return nil
	}
	return o.DealIds
}

func (o *CrmContact) GetEmails() []CrmEmail {
	if o == nil {
		return nil
	}
	return o.Emails
}

func (o *CrmContact) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CrmContact) GetLinkUrls() []string {
	if o == nil {
		return nil
	}
	return o.LinkUrls
}

func (o *CrmContact) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CrmContact) GetRaw() *CrmContactRaw {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *CrmContact) GetTelephones() []CrmTelephone {
	if o == nil {
		return nil
	}
	return o.Telephones
}

func (o *CrmContact) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *CrmContact) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *CrmContact) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

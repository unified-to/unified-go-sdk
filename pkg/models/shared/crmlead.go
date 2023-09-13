// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

type CrmLead struct {
	Active        *bool                   `json:"active,omitempty"`
	Address       *PropertyCrmLeadAddress `json:"address,omitempty"`
	CompanyID     *string                 `json:"company_id,omitempty"`
	ContactID     *string                 `json:"contact_id,omitempty"`
	CreatedAt     *types.Date             `json:"created_at,omitempty"`
	CreatorUserID *string                 `json:"creator_user_id,omitempty"`
	Emails        []CrmEmail              `json:"emails,omitempty"`
	ID            *string                 `json:"id,omitempty"`
	Name          *string                 `json:"name,omitempty"`
	Raw           *PropertyCrmLeadRaw     `json:"raw,omitempty"`
	Telephones    []CrmTelephone          `json:"telephones,omitempty"`
	UpdatedAt     *types.Date             `json:"updated_at,omitempty"`
	UserID        *string                 `json:"user_id,omitempty"`
}

func (o *CrmLead) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *CrmLead) GetAddress() *PropertyCrmLeadAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *CrmLead) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *CrmLead) GetContactID() *string {
	if o == nil {
		return nil
	}
	return o.ContactID
}

func (o *CrmLead) GetCreatedAt() *types.Date {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CrmLead) GetCreatorUserID() *string {
	if o == nil {
		return nil
	}
	return o.CreatorUserID
}

func (o *CrmLead) GetEmails() []CrmEmail {
	if o == nil {
		return nil
	}
	return o.Emails
}

func (o *CrmLead) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CrmLead) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CrmLead) GetRaw() *PropertyCrmLeadRaw {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *CrmLead) GetTelephones() []CrmTelephone {
	if o == nil {
		return nil
	}
	return o.Telephones
}

func (o *CrmLead) GetUpdatedAt() *types.Date {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *CrmLead) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

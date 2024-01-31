// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type CrmLead struct {
	Address       *PropertyCrmLeadAddress `json:"address,omitempty"`
	CompanyID     *string                 `json:"company_id,omitempty"`
	CompanyName   *string                 `json:"company_name,omitempty"`
	ContactID     *string                 `json:"contact_id,omitempty"`
	CreatedAt     *time.Time              `json:"created_at,omitempty"`
	CreatorUserID *string                 `json:"creator_user_id,omitempty"`
	Emails        []CrmEmail              `json:"emails,omitempty"`
	ID            *string                 `json:"id,omitempty"`
	IsActive      *bool                   `json:"is_active,omitempty"`
	Name          *string                 `json:"name,omitempty"`
	Raw           map[string]interface{}  `json:"raw,omitempty"`
	Source        *string                 `json:"source,omitempty"`
	Status        *string                 `json:"status,omitempty"`
	Telephones    []CrmTelephone          `json:"telephones,omitempty"`
	UpdatedAt     *time.Time              `json:"updated_at,omitempty"`
	UserID        *string                 `json:"user_id,omitempty"`
}

func (c CrmLead) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CrmLead) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
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

func (o *CrmLead) GetCompanyName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyName
}

func (o *CrmLead) GetContactID() *string {
	if o == nil {
		return nil
	}
	return o.ContactID
}

func (o *CrmLead) GetCreatedAt() *time.Time {
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

func (o *CrmLead) GetIsActive() *bool {
	if o == nil {
		return nil
	}
	return o.IsActive
}

func (o *CrmLead) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CrmLead) GetRaw() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *CrmLead) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *CrmLead) GetStatus() *string {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *CrmLead) GetTelephones() []CrmTelephone {
	if o == nil {
		return nil
	}
	return o.Telephones
}

func (o *CrmLead) GetUpdatedAt() *time.Time {
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

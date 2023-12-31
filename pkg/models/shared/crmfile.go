// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type CrmFile struct {
	ActivityID  *string             `json:"activity_id,omitempty"`
	CompanyID   *string             `json:"company_id,omitempty"`
	ContactID   *string             `json:"contact_id,omitempty"`
	CreatedAt   *time.Time          `json:"created_at,omitempty"`
	DealID      *string             `json:"deal_id,omitempty"`
	Description *string             `json:"description,omitempty"`
	FileName    *string             `json:"file_name,omitempty"`
	FileSize    *float64            `json:"file_size,omitempty"`
	FileType    *string             `json:"file_type,omitempty"`
	FileURL     *string             `json:"file_url,omitempty"`
	ID          *string             `json:"id,omitempty"`
	IsActive    *bool               `json:"is_active,omitempty"`
	LeadID      *string             `json:"lead_id,omitempty"`
	Raw         *PropertyCrmFileRaw `json:"raw,omitempty"`
	UpdatedAt   *time.Time          `json:"updated_at,omitempty"`
	UserID      *string             `json:"user_id,omitempty"`
}

func (c CrmFile) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CrmFile) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CrmFile) GetActivityID() *string {
	if o == nil {
		return nil
	}
	return o.ActivityID
}

func (o *CrmFile) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *CrmFile) GetContactID() *string {
	if o == nil {
		return nil
	}
	return o.ContactID
}

func (o *CrmFile) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CrmFile) GetDealID() *string {
	if o == nil {
		return nil
	}
	return o.DealID
}

func (o *CrmFile) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CrmFile) GetFileName() *string {
	if o == nil {
		return nil
	}
	return o.FileName
}

func (o *CrmFile) GetFileSize() *float64 {
	if o == nil {
		return nil
	}
	return o.FileSize
}

func (o *CrmFile) GetFileType() *string {
	if o == nil {
		return nil
	}
	return o.FileType
}

func (o *CrmFile) GetFileURL() *string {
	if o == nil {
		return nil
	}
	return o.FileURL
}

func (o *CrmFile) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CrmFile) GetIsActive() *bool {
	if o == nil {
		return nil
	}
	return o.IsActive
}

func (o *CrmFile) GetLeadID() *string {
	if o == nil {
		return nil
	}
	return o.LeadID
}

func (o *CrmFile) GetRaw() *PropertyCrmFileRaw {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *CrmFile) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *CrmFile) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

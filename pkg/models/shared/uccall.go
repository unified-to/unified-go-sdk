// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type UcCall struct {
	AgencyID  string     `json:"agency_id"`
	ContactID *string    `json:"contact_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	EndAt     *time.Time `json:"end_at,omitempty"`
	ID        *string    `json:"id,omitempty"`
	// The raw data returned by the integration for this call
	Raw     *PropertyUcCallRaw `json:"raw,omitempty"`
	StartAt *time.Time         `json:"start_at,omitempty"`
	// The telephone number called
	Telephone *PropertyUcCallTelephone `json:"telephone,omitempty"`
	UpdatedAt *time.Time               `json:"updated_at,omitempty"`
}

func (u UcCall) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UcCall) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UcCall) GetAgencyID() string {
	if o == nil {
		return ""
	}
	return o.AgencyID
}

func (o *UcCall) GetContactID() *string {
	if o == nil {
		return nil
	}
	return o.ContactID
}

func (o *UcCall) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UcCall) GetEndAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndAt
}

func (o *UcCall) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UcCall) GetRaw() *PropertyUcCallRaw {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *UcCall) GetStartAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartAt
}

func (o *UcCall) GetTelephone() *PropertyUcCallTelephone {
	if o == nil {
		return nil
	}
	return o.Telephone
}

func (o *UcCall) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

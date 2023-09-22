// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type CrmTeam struct {
	CreatedAt   *time.Time          `json:"created_at,omitempty"`
	Description *string             `json:"description,omitempty"`
	ID          *string             `json:"id,omitempty"`
	Name        *string             `json:"name,omitempty"`
	Raw         *PropertyCrmTeamRaw `json:"raw,omitempty"`
	UpdatedAt   *time.Time          `json:"updated_at,omitempty"`
	UserIds     []string            `json:"user_ids,omitempty"`
}

func (c CrmTeam) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CrmTeam) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CrmTeam) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CrmTeam) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CrmTeam) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CrmTeam) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CrmTeam) GetRaw() *PropertyCrmTeamRaw {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *CrmTeam) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *CrmTeam) GetUserIds() []string {
	if o == nil {
		return nil
	}
	return o.UserIds
}

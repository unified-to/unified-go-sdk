// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type RepoOrganization struct {
	AvatarURL   *string        `json:"avatar_url,omitempty"`
	CreatedAt   *time.Time     `json:"created_at,omitempty"`
	Description *string        `json:"description,omitempty"`
	ID          *string        `json:"id,omitempty"`
	Name        *string        `json:"name,omitempty"`
	Raw         map[string]any `json:"raw,omitempty"`
	UpdatedAt   *time.Time     `json:"updated_at,omitempty"`
	// id values of the users/employees associated with this organization
	UserIds []string `json:"user_ids,omitempty"`
	WebURL  *string  `json:"web_url,omitempty"`
}

func (r RepoOrganization) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RepoOrganization) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RepoOrganization) GetAvatarURL() *string {
	if o == nil {
		return nil
	}
	return o.AvatarURL
}

func (o *RepoOrganization) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *RepoOrganization) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *RepoOrganization) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RepoOrganization) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *RepoOrganization) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *RepoOrganization) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *RepoOrganization) GetUserIds() []string {
	if o == nil {
		return nil
	}
	return o.UserIds
}

func (o *RepoOrganization) GetWebURL() *string {
	if o == nil {
		return nil
	}
	return o.WebURL
}

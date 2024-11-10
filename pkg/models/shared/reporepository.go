// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type RepoRepository struct {
	CreatedAt   *time.Time     `json:"created_at,omitempty"`
	Description *string        `json:"description,omitempty"`
	ID          *string        `json:"id,omitempty"`
	IsPrivate   *bool          `json:"is_private,omitempty"`
	Name        string         `json:"name"`
	OrgID       *string        `json:"org_id,omitempty"`
	Owner       *string        `json:"owner,omitempty"`
	Raw         map[string]any `json:"raw,omitempty"`
	UpdatedAt   *time.Time     `json:"updated_at,omitempty"`
	WebURL      *string        `json:"web_url,omitempty"`
}

func (r RepoRepository) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RepoRepository) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RepoRepository) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *RepoRepository) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *RepoRepository) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RepoRepository) GetIsPrivate() *bool {
	if o == nil {
		return nil
	}
	return o.IsPrivate
}

func (o *RepoRepository) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *RepoRepository) GetOrgID() *string {
	if o == nil {
		return nil
	}
	return o.OrgID
}

func (o *RepoRepository) GetOwner() *string {
	if o == nil {
		return nil
	}
	return o.Owner
}

func (o *RepoRepository) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *RepoRepository) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *RepoRepository) GetWebURL() *string {
	if o == nil {
		return nil
	}
	return o.WebURL
}

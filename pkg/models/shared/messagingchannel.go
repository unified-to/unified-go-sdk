// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type MessagingChannel struct {
	CreatedAt       *time.Time        `json:"created_at,omitempty"`
	Description     *string           `json:"description,omitempty"`
	HasSubchannels  *bool             `json:"has_subchannels,omitempty"`
	ID              *string           `json:"id,omitempty"`
	IsActive        *bool             `json:"is_active,omitempty"`
	IsPrivate       *bool             `json:"is_private,omitempty"`
	Members         []MessagingMember `json:"members,omitempty"`
	Name            string            `json:"name"`
	ParentChannelID *string           `json:"parent_channel_id,omitempty"`
	Raw             map[string]any    `json:"raw,omitempty"`
	UpdatedAt       *string           `json:"updated_at,omitempty"`
	WebURL          *string           `json:"web_url,omitempty"`
}

func (m MessagingChannel) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *MessagingChannel) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *MessagingChannel) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *MessagingChannel) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *MessagingChannel) GetHasSubchannels() *bool {
	if o == nil {
		return nil
	}
	return o.HasSubchannels
}

func (o *MessagingChannel) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *MessagingChannel) GetIsActive() *bool {
	if o == nil {
		return nil
	}
	return o.IsActive
}

func (o *MessagingChannel) GetIsPrivate() *bool {
	if o == nil {
		return nil
	}
	return o.IsPrivate
}

func (o *MessagingChannel) GetMembers() []MessagingMember {
	if o == nil {
		return nil
	}
	return o.Members
}

func (o *MessagingChannel) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *MessagingChannel) GetParentChannelID() *string {
	if o == nil {
		return nil
	}
	return o.ParentChannelID
}

func (o *MessagingChannel) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *MessagingChannel) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *MessagingChannel) GetWebURL() *string {
	if o == nil {
		return nil
	}
	return o.WebURL
}

// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type UcRecording struct {
	CallID       *string            `json:"call_id,omitempty"`
	ContactID    *string            `json:"contact_id,omitempty"`
	ContactName  *string            `json:"contact_name,omitempty"`
	ContactPhone *string            `json:"contact_phone,omitempty"`
	CreatedAt    *time.Time         `json:"created_at,omitempty"`
	EndAt        *time.Time         `json:"end_at,omitempty"`
	ExpiresAt    *time.Time         `json:"expires_at,omitempty"`
	ID           *string            `json:"id,omitempty"`
	Media        []UcRecordingMedia `json:"media,omitempty"`
	Raw          map[string]any     `json:"raw,omitempty"`
	StartAt      *time.Time         `json:"start_at,omitempty"`
	UpdatedAt    *time.Time         `json:"updated_at,omitempty"`
	UserID       *string            `json:"user_id,omitempty"`
	WebURL       *string            `json:"web_url,omitempty"`
}

func (u UcRecording) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UcRecording) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UcRecording) GetCallID() *string {
	if o == nil {
		return nil
	}
	return o.CallID
}

func (o *UcRecording) GetContactID() *string {
	if o == nil {
		return nil
	}
	return o.ContactID
}

func (o *UcRecording) GetContactName() *string {
	if o == nil {
		return nil
	}
	return o.ContactName
}

func (o *UcRecording) GetContactPhone() *string {
	if o == nil {
		return nil
	}
	return o.ContactPhone
}

func (o *UcRecording) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *UcRecording) GetEndAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndAt
}

func (o *UcRecording) GetExpiresAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *UcRecording) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *UcRecording) GetMedia() []UcRecordingMedia {
	if o == nil {
		return nil
	}
	return o.Media
}

func (o *UcRecording) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *UcRecording) GetStartAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartAt
}

func (o *UcRecording) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *UcRecording) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *UcRecording) GetWebURL() *string {
	if o == nil {
		return nil
	}
	return o.WebURL
}

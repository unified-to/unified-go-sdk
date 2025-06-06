// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type CalendarRecording struct {
	CreatedAt *time.Time               `json:"created_at,omitempty"`
	EndAt     *time.Time               `json:"end_at,omitempty"`
	EventID   *string                  `json:"event_id,omitempty"`
	ExpiresAt *time.Time               `json:"expires_at,omitempty"`
	ID        *string                  `json:"id,omitempty"`
	Media     []CalendarRecordingMedia `json:"media,omitempty"`
	Raw       map[string]any           `json:"raw,omitempty"`
	StartAt   *time.Time               `json:"start_at,omitempty"`
	UpdatedAt *time.Time               `json:"updated_at,omitempty"`
	WebURL    *string                  `json:"web_url,omitempty"`
}

func (c CalendarRecording) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CalendarRecording) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CalendarRecording) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CalendarRecording) GetEndAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndAt
}

func (o *CalendarRecording) GetEventID() *string {
	if o == nil {
		return nil
	}
	return o.EventID
}

func (o *CalendarRecording) GetExpiresAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ExpiresAt
}

func (o *CalendarRecording) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CalendarRecording) GetMedia() []CalendarRecordingMedia {
	if o == nil {
		return nil
	}
	return o.Media
}

func (o *CalendarRecording) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *CalendarRecording) GetStartAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartAt
}

func (o *CalendarRecording) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *CalendarRecording) GetWebURL() *string {
	if o == nil {
		return nil
	}
	return o.WebURL
}

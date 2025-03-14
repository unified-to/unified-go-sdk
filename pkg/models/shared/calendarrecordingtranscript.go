// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type CalendarRecordingTranscript struct {
	Attendee *PropertyCalendarRecordingTranscriptAttendee `json:"attendee,omitempty"`
	EndAt    *time.Time                                   `json:"end_at,omitempty"`
	Language *string                                      `json:"language,omitempty"`
	StartAt  *time.Time                                   `json:"start_at,omitempty"`
	Text     string                                       `json:"text"`
}

func (c CalendarRecordingTranscript) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CalendarRecordingTranscript) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CalendarRecordingTranscript) GetAttendee() *PropertyCalendarRecordingTranscriptAttendee {
	if o == nil {
		return nil
	}
	return o.Attendee
}

func (o *CalendarRecordingTranscript) GetEndAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndAt
}

func (o *CalendarRecordingTranscript) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *CalendarRecordingTranscript) GetStartAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartAt
}

func (o *CalendarRecordingTranscript) GetText() string {
	if o == nil {
		return ""
	}
	return o.Text
}

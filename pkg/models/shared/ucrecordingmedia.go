// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type UcRecordingMedia struct {
	EndAt                 *time.Time              `json:"end_at,omitempty"`
	Language              *string                 `json:"language,omitempty"`
	RecordingDownloadURL  *string                 `json:"recording_download_url,omitempty"`
	StartAt               *time.Time              `json:"start_at,omitempty"`
	TranscriptDownloadURL *string                 `json:"transcript_download_url,omitempty"`
	Transcripts           []UcRecordingTranscript `json:"transcripts,omitempty"`
}

func (u UcRecordingMedia) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UcRecordingMedia) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UcRecordingMedia) GetEndAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndAt
}

func (o *UcRecordingMedia) GetLanguage() *string {
	if o == nil {
		return nil
	}
	return o.Language
}

func (o *UcRecordingMedia) GetRecordingDownloadURL() *string {
	if o == nil {
		return nil
	}
	return o.RecordingDownloadURL
}

func (o *UcRecordingMedia) GetStartAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartAt
}

func (o *UcRecordingMedia) GetTranscriptDownloadURL() *string {
	if o == nil {
		return nil
	}
	return o.TranscriptDownloadURL
}

func (o *UcRecordingMedia) GetTranscripts() []UcRecordingTranscript {
	if o == nil {
		return nil
	}
	return o.Transcripts
}

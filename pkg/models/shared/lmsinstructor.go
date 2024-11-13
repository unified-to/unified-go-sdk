// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type LmsInstructor struct {
	CreatedAt  *time.Time     `json:"created_at,omitempty"`
	Emails     []LmsEmail     `json:"emails,omitempty"`
	ID         *string        `json:"id,omitempty"`
	ImageURL   *string        `json:"image_url,omitempty"`
	Name       *string        `json:"name,omitempty"`
	Raw        map[string]any `json:"raw,omitempty"`
	Telephones []LmsTelephone `json:"telephones,omitempty"`
	Title      *string        `json:"title,omitempty"`
	UpdatedAt  *time.Time     `json:"updated_at,omitempty"`
}

func (l LmsInstructor) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *LmsInstructor) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *LmsInstructor) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *LmsInstructor) GetEmails() []LmsEmail {
	if o == nil {
		return nil
	}
	return o.Emails
}

func (o *LmsInstructor) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *LmsInstructor) GetImageURL() *string {
	if o == nil {
		return nil
	}
	return o.ImageURL
}

func (o *LmsInstructor) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *LmsInstructor) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *LmsInstructor) GetTelephones() []LmsTelephone {
	if o == nil {
		return nil
	}
	return o.Telephones
}

func (o *LmsInstructor) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *LmsInstructor) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
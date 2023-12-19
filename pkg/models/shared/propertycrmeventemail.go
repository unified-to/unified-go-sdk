// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

// PropertyCrmEventEmail - The email object, when type = email
type PropertyCrmEventEmail struct {
	Body *string `json:"body,omitempty"`
	// The event email's cc name & email (name <test@test.com>)
	Cc        []string   `json:"cc,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	From      *string    `json:"from,omitempty"`
	Subject   *string    `json:"subject,omitempty"`
	// The event email's to name & email (name <test@test.com>)
	To []string `json:"to,omitempty"`
}

func (p PropertyCrmEventEmail) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PropertyCrmEventEmail) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PropertyCrmEventEmail) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *PropertyCrmEventEmail) GetCc() []string {
	if o == nil {
		return nil
	}
	return o.Cc
}

func (o *PropertyCrmEventEmail) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *PropertyCrmEventEmail) GetFrom() *string {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *PropertyCrmEventEmail) GetSubject() *string {
	if o == nil {
		return nil
	}
	return o.Subject
}

func (o *PropertyCrmEventEmail) GetTo() []string {
	if o == nil {
		return nil
	}
	return o.To
}

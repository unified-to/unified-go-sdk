// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type Origin string

const (
	OriginAgency     Origin = "AGENCY"
	OriginApplied    Origin = "APPLIED"
	OriginInternal   Origin = "INTERNAL"
	OriginReferred   Origin = "REFERRED"
	OriginSourced    Origin = "SOURCED"
	OriginUniversity Origin = "UNIVERSITY"
)

func (e Origin) ToPointer() *Origin {
	return &e
}

func (e *Origin) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "AGENCY":
		fallthrough
	case "APPLIED":
		fallthrough
	case "INTERNAL":
		fallthrough
	case "REFERRED":
		fallthrough
	case "SOURCED":
		fallthrough
	case "UNIVERSITY":
		*e = Origin(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Origin: %v", v)
	}
}

type AtsCandidate struct {
	Address            *PropertyAtsCandidateAddress `json:"address,omitempty"`
	CompanyID          *string                      `json:"company_id,omitempty"`
	CompanyName        *string                      `json:"company_name,omitempty"`
	CreatedAt          *time.Time                   `json:"created_at,omitempty"`
	DateOfBirth        *time.Time                   `json:"date_of_birth,omitempty"`
	Emails             []AtsEmail                   `json:"emails,omitempty"`
	ExternalIdentifier *string                      `json:"external_identifier,omitempty"`
	ID                 *string                      `json:"id,omitempty"`
	ImageURL           *string                      `json:"image_url,omitempty"`
	// a list of social media links associated with the candidate. eg. LinkedIn URL
	LinkUrls   []string       `json:"link_urls,omitempty"`
	Name       *string        `json:"name,omitempty"`
	Origin     *Origin        `json:"origin,omitempty"`
	Raw        map[string]any `json:"raw,omitempty"`
	Sources    []string       `json:"sources,omitempty"`
	Tags       []string       `json:"tags,omitempty"`
	Telephones []AtsTelephone `json:"telephones,omitempty"`
	Title      *string        `json:"title,omitempty"`
	UpdatedAt  *time.Time     `json:"updated_at,omitempty"`
	UserID     *string        `json:"user_id,omitempty"`
}

func (a AtsCandidate) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AtsCandidate) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AtsCandidate) GetAddress() *PropertyAtsCandidateAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *AtsCandidate) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *AtsCandidate) GetCompanyName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyName
}

func (o *AtsCandidate) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *AtsCandidate) GetDateOfBirth() *time.Time {
	if o == nil {
		return nil
	}
	return o.DateOfBirth
}

func (o *AtsCandidate) GetEmails() []AtsEmail {
	if o == nil {
		return nil
	}
	return o.Emails
}

func (o *AtsCandidate) GetExternalIdentifier() *string {
	if o == nil {
		return nil
	}
	return o.ExternalIdentifier
}

func (o *AtsCandidate) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *AtsCandidate) GetImageURL() *string {
	if o == nil {
		return nil
	}
	return o.ImageURL
}

func (o *AtsCandidate) GetLinkUrls() []string {
	if o == nil {
		return nil
	}
	return o.LinkUrls
}

func (o *AtsCandidate) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *AtsCandidate) GetOrigin() *Origin {
	if o == nil {
		return nil
	}
	return o.Origin
}

func (o *AtsCandidate) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *AtsCandidate) GetSources() []string {
	if o == nil {
		return nil
	}
	return o.Sources
}

func (o *AtsCandidate) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *AtsCandidate) GetTelephones() []AtsTelephone {
	if o == nil {
		return nil
	}
	return o.Telephones
}

func (o *AtsCandidate) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *AtsCandidate) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *AtsCandidate) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

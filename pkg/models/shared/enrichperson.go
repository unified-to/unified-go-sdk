// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"time"
)

type EnrichPersonGender string

const (
	EnrichPersonGenderMale   EnrichPersonGender = "MALE"
	EnrichPersonGenderFemale EnrichPersonGender = "FEMALE"
)

func (e EnrichPersonGender) ToPointer() *EnrichPersonGender {
	return &e
}

func (e *EnrichPersonGender) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MALE":
		fallthrough
	case "FEMALE":
		*e = EnrichPersonGender(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EnrichPersonGender: %v", v)
	}
}

// EnrichPerson - A person object from an enrichment integration
type EnrichPerson struct {
	// The address of the person
	Address       *PropertyEnrichPersonAddress `json:"address,omitempty"`
	Bio           *string                      `json:"bio,omitempty"`
	Birthdate     *string                      `json:"birthdate,omitempty"`
	Company       *string                      `json:"company,omitempty"`
	CompanyDomain *string                      `json:"company_domain,omitempty"`
	CreatedAt     *time.Time                   `json:"created_at,omitempty"`
	// An array of email addresses for this person
	Emails         []EnrichEmail       `json:"emails,omitempty"`
	FacebookURL    *string             `json:"facebook_url,omitempty"`
	Gender         *EnrichPersonGender `json:"gender,omitempty"`
	GithubURL      *string             `json:"github_url,omitempty"`
	GithubUsername *string             `json:"github_username,omitempty"`
	ID             *string             `json:"id,omitempty"`
	ImageURL       *string             `json:"image_url,omitempty"`
	LinkedinURL    *string             `json:"linkedin_url,omitempty"`
	Name           *string             `json:"name,omitempty"`
	// The raw data returned by the integration for this person
	Raw PropertyEnrichPersonRaw `json:"raw"`
	// An array of telephones for this person
	Telephones    []EnrichTelephone         `json:"telephones,omitempty"`
	Timezone      *string                   `json:"timezone,omitempty"`
	Title         *string                   `json:"title,omitempty"`
	TwitterHandle *string                   `json:"twitter_handle,omitempty"`
	TwitterURL    *string                   `json:"twitter_url,omitempty"`
	UpdatedAt     *time.Time                `json:"updated_at,omitempty"`
	UtcOffset     *float64                  `json:"utc_offset,omitempty"`
	WorkHistories []EnrichPersonWorkHistory `json:"work_histories,omitempty"`
}

func (o *EnrichPerson) GetAddress() *PropertyEnrichPersonAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *EnrichPerson) GetBio() *string {
	if o == nil {
		return nil
	}
	return o.Bio
}

func (o *EnrichPerson) GetBirthdate() *string {
	if o == nil {
		return nil
	}
	return o.Birthdate
}

func (o *EnrichPerson) GetCompany() *string {
	if o == nil {
		return nil
	}
	return o.Company
}

func (o *EnrichPerson) GetCompanyDomain() *string {
	if o == nil {
		return nil
	}
	return o.CompanyDomain
}

func (o *EnrichPerson) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *EnrichPerson) GetEmails() []EnrichEmail {
	if o == nil {
		return nil
	}
	return o.Emails
}

func (o *EnrichPerson) GetFacebookURL() *string {
	if o == nil {
		return nil
	}
	return o.FacebookURL
}

func (o *EnrichPerson) GetGender() *EnrichPersonGender {
	if o == nil {
		return nil
	}
	return o.Gender
}

func (o *EnrichPerson) GetGithubURL() *string {
	if o == nil {
		return nil
	}
	return o.GithubURL
}

func (o *EnrichPerson) GetGithubUsername() *string {
	if o == nil {
		return nil
	}
	return o.GithubUsername
}

func (o *EnrichPerson) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *EnrichPerson) GetImageURL() *string {
	if o == nil {
		return nil
	}
	return o.ImageURL
}

func (o *EnrichPerson) GetLinkedinURL() *string {
	if o == nil {
		return nil
	}
	return o.LinkedinURL
}

func (o *EnrichPerson) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *EnrichPerson) GetRaw() PropertyEnrichPersonRaw {
	if o == nil {
		return PropertyEnrichPersonRaw{}
	}
	return o.Raw
}

func (o *EnrichPerson) GetTelephones() []EnrichTelephone {
	if o == nil {
		return nil
	}
	return o.Telephones
}

func (o *EnrichPerson) GetTimezone() *string {
	if o == nil {
		return nil
	}
	return o.Timezone
}

func (o *EnrichPerson) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *EnrichPerson) GetTwitterHandle() *string {
	if o == nil {
		return nil
	}
	return o.TwitterHandle
}

func (o *EnrichPerson) GetTwitterURL() *string {
	if o == nil {
		return nil
	}
	return o.TwitterURL
}

func (o *EnrichPerson) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *EnrichPerson) GetUtcOffset() *float64 {
	if o == nil {
		return nil
	}
	return o.UtcOffset
}

func (o *EnrichPerson) GetWorkHistories() []EnrichPersonWorkHistory {
	if o == nil {
		return nil
	}
	return o.WorkHistories
}

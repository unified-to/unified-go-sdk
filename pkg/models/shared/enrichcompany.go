// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

// EnrichCompany - A company object from an enrichment integration
type EnrichCompany struct {
	// The address of the company
	Address       *PropertyEnrichCompanyAddress `json:"address,omitempty"`
	AlexaRank     *float64                      `json:"alexa_rank,omitempty"`
	CreatedAt     *time.Time                    `json:"created_at,omitempty"`
	CrunchbaseURL *string                       `json:"crunchbase_url,omitempty"`
	Description   *string                       `json:"description,omitempty"`
	Domain        *string                       `json:"domain,omitempty"`
	Employees     *string                       `json:"employees,omitempty"`
	Exchange      *string                       `json:"exchange,omitempty"`
	FacebookURL   *string                       `json:"facebook_url,omitempty"`
	ID            *string                       `json:"id,omitempty"`
	Industry      *string                       `json:"industry,omitempty"`
	InstagramURL  *string                       `json:"instagram_url,omitempty"`
	LinkedinURL   *string                       `json:"linkedin_url,omitempty"`
	LogoURL       *string                       `json:"logo_url,omitempty"`
	NaicsCode     *float64                      `json:"naics_code,omitempty"`
	Name          *string                       `json:"name,omitempty"`
	// The raw data returned by the integration for this company
	Raw     map[string]any `json:"raw,omitempty"`
	Revenue *string        `json:"revenue,omitempty"`
	SicCode *float64       `json:"sic_code,omitempty"`
	Stock   *string        `json:"stock,omitempty"`
	// An array of telephones for this company
	Telephones    []EnrichTelephone `json:"telephones,omitempty"`
	TwitterHandle *string           `json:"twitter_handle,omitempty"`
	TwitterURL    *string           `json:"twitter_url,omitempty"`
	UpdatedAt     *time.Time        `json:"updated_at,omitempty"`
	YearFounded   *float64          `json:"year_founded,omitempty"`
	YelpURL       *string           `json:"yelp_url,omitempty"`
	YoutubeURL    *string           `json:"youtube_url,omitempty"`
}

func (e EnrichCompany) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(e, "", false)
}

func (e *EnrichCompany) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &e, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *EnrichCompany) GetAddress() *PropertyEnrichCompanyAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *EnrichCompany) GetAlexaRank() *float64 {
	if o == nil {
		return nil
	}
	return o.AlexaRank
}

func (o *EnrichCompany) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *EnrichCompany) GetCrunchbaseURL() *string {
	if o == nil {
		return nil
	}
	return o.CrunchbaseURL
}

func (o *EnrichCompany) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *EnrichCompany) GetDomain() *string {
	if o == nil {
		return nil
	}
	return o.Domain
}

func (o *EnrichCompany) GetEmployees() *string {
	if o == nil {
		return nil
	}
	return o.Employees
}

func (o *EnrichCompany) GetExchange() *string {
	if o == nil {
		return nil
	}
	return o.Exchange
}

func (o *EnrichCompany) GetFacebookURL() *string {
	if o == nil {
		return nil
	}
	return o.FacebookURL
}

func (o *EnrichCompany) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *EnrichCompany) GetIndustry() *string {
	if o == nil {
		return nil
	}
	return o.Industry
}

func (o *EnrichCompany) GetInstagramURL() *string {
	if o == nil {
		return nil
	}
	return o.InstagramURL
}

func (o *EnrichCompany) GetLinkedinURL() *string {
	if o == nil {
		return nil
	}
	return o.LinkedinURL
}

func (o *EnrichCompany) GetLogoURL() *string {
	if o == nil {
		return nil
	}
	return o.LogoURL
}

func (o *EnrichCompany) GetNaicsCode() *float64 {
	if o == nil {
		return nil
	}
	return o.NaicsCode
}

func (o *EnrichCompany) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *EnrichCompany) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *EnrichCompany) GetRevenue() *string {
	if o == nil {
		return nil
	}
	return o.Revenue
}

func (o *EnrichCompany) GetSicCode() *float64 {
	if o == nil {
		return nil
	}
	return o.SicCode
}

func (o *EnrichCompany) GetStock() *string {
	if o == nil {
		return nil
	}
	return o.Stock
}

func (o *EnrichCompany) GetTelephones() []EnrichTelephone {
	if o == nil {
		return nil
	}
	return o.Telephones
}

func (o *EnrichCompany) GetTwitterHandle() *string {
	if o == nil {
		return nil
	}
	return o.TwitterHandle
}

func (o *EnrichCompany) GetTwitterURL() *string {
	if o == nil {
		return nil
	}
	return o.TwitterURL
}

func (o *EnrichCompany) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *EnrichCompany) GetYearFounded() *float64 {
	if o == nil {
		return nil
	}
	return o.YearFounded
}

func (o *EnrichCompany) GetYelpURL() *string {
	if o == nil {
		return nil
	}
	return o.YelpURL
}

func (o *EnrichCompany) GetYoutubeURL() *string {
	if o == nil {
		return nil
	}
	return o.YoutubeURL
}

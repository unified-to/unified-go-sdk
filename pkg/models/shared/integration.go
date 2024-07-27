// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

// Integration - Informational object for supported integrations.
type Integration struct {
	ActiveHealthyConnections *float64       `json:"active_healthy_connections,omitempty"`
	API                      map[string]any `json:"api,omitempty"`
	APIDocsURL               *string        `json:"api_docs_url,omitempty"`
	Beta                     *bool          `json:"beta,omitempty"`
	// The categories of support solutions that this integration has
	Categories           []PropertyIntegrationCategories `json:"categories"`
	Color                *string                         `json:"color,omitempty"`
	CreatedAt            *string                         `json:"created_at,omitempty"`
	Description          *string                         `json:"description,omitempty"`
	FaIcon               *string                         `json:"fa_icon,omitempty"`
	Featured             *bool                           `json:"featured,omitempty"`
	InProgress           bool                            `json:"in_progress"`
	IsActive             *bool                           `json:"is_active,omitempty"`
	LogoURL              *string                         `json:"logo_url,omitempty"`
	Name                 string                          `json:"name"`
	Popularity           *float64                        `json:"popularity,omitempty"`
	RateLimitDescription *string                         `json:"rate_limit_description,omitempty"`
	Support              map[string]IntegrationSupport   `json:"support,omitempty"`
	TestedAt             *time.Time                      `json:"tested_at,omitempty"`
	TextColor            *string                         `json:"text_color,omitempty"`
	// instructions for the user on how to find the token/key
	TokenInstructions []string `json:"token_instructions,omitempty"`
	// if auth_types = 'token'
	TokenNames []string `json:"token_names,omitempty"`
	Type       string   `json:"type"`
	UpdatedAt  *string  `json:"updated_at,omitempty"`
	WebURL     *string  `json:"web_url,omitempty"`
}

func (i Integration) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(i, "", false)
}

func (i *Integration) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &i, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Integration) GetActiveHealthyConnections() *float64 {
	if o == nil {
		return nil
	}
	return o.ActiveHealthyConnections
}

func (o *Integration) GetAPI() map[string]any {
	if o == nil {
		return nil
	}
	return o.API
}

func (o *Integration) GetAPIDocsURL() *string {
	if o == nil {
		return nil
	}
	return o.APIDocsURL
}

func (o *Integration) GetBeta() *bool {
	if o == nil {
		return nil
	}
	return o.Beta
}

func (o *Integration) GetCategories() []PropertyIntegrationCategories {
	if o == nil {
		return []PropertyIntegrationCategories{}
	}
	return o.Categories
}

func (o *Integration) GetColor() *string {
	if o == nil {
		return nil
	}
	return o.Color
}

func (o *Integration) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Integration) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *Integration) GetFaIcon() *string {
	if o == nil {
		return nil
	}
	return o.FaIcon
}

func (o *Integration) GetFeatured() *bool {
	if o == nil {
		return nil
	}
	return o.Featured
}

func (o *Integration) GetInProgress() bool {
	if o == nil {
		return false
	}
	return o.InProgress
}

func (o *Integration) GetIsActive() *bool {
	if o == nil {
		return nil
	}
	return o.IsActive
}

func (o *Integration) GetLogoURL() *string {
	if o == nil {
		return nil
	}
	return o.LogoURL
}

func (o *Integration) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Integration) GetPopularity() *float64 {
	if o == nil {
		return nil
	}
	return o.Popularity
}

func (o *Integration) GetRateLimitDescription() *string {
	if o == nil {
		return nil
	}
	return o.RateLimitDescription
}

func (o *Integration) GetSupport() map[string]IntegrationSupport {
	if o == nil {
		return nil
	}
	return o.Support
}

func (o *Integration) GetTestedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.TestedAt
}

func (o *Integration) GetTextColor() *string {
	if o == nil {
		return nil
	}
	return o.TextColor
}

func (o *Integration) GetTokenInstructions() []string {
	if o == nil {
		return nil
	}
	return o.TokenInstructions
}

func (o *Integration) GetTokenNames() []string {
	if o == nil {
		return nil
	}
	return o.TokenNames
}

func (o *Integration) GetType() string {
	if o == nil {
		return ""
	}
	return o.Type
}

func (o *Integration) GetUpdatedAt() *string {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *Integration) GetWebURL() *string {
	if o == nil {
		return nil
	}
	return o.WebURL
}

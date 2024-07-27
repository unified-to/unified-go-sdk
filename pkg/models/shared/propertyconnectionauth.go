// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

// PropertyConnectionAuth - An authentication object that represents a specific authorized user's connection to an integration.
type PropertyConnectionAuth struct {
	AccessToken    *string        `json:"access_token,omitempty"`
	APIURL         *string        `json:"api_url,omitempty"`
	AppID          *string        `json:"app_id,omitempty"`
	AuthorizeURL   *string        `json:"authorize_url,omitempty"`
	ClientID       *string        `json:"client_id,omitempty"`
	ClientSecret   *string        `json:"client_secret,omitempty"`
	ConsumerKey    *string        `json:"consumer_key,omitempty"`
	ConsumerSecret *string        `json:"consumer_secret,omitempty"`
	DevAPIKey      *string        `json:"dev_api_key,omitempty"`
	Emails         []string       `json:"emails,omitempty"`
	ExpiresIn      *float64       `json:"expires_in,omitempty"`
	ExpiryDate     *time.Time     `json:"expiry_date,omitempty"`
	Key            *string        `json:"key,omitempty"`
	Meta           map[string]any `json:"meta,omitempty"`
	Name           *string        `json:"name,omitempty"`
	// When integration.auth_type = "other", this field contains the authentication credentials in the same order as token_names
	OtherAuthInfo           []string   `json:"other_auth_info,omitempty"`
	Pem                     *string    `json:"pem,omitempty"`
	RefreshToken            *string    `json:"refresh_token,omitempty"`
	RefreshTokenExpiresDate *time.Time `json:"refresh_token_expires_date,omitempty"`
	RefreshTokenExpiresIn   *float64   `json:"refresh_token_expires_in,omitempty"`
	State                   *string    `json:"state,omitempty"`
	Token                   *string    `json:"token,omitempty"`
	TokenURL                *string    `json:"token_url,omitempty"`
}

func (p PropertyConnectionAuth) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PropertyConnectionAuth) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PropertyConnectionAuth) GetAccessToken() *string {
	if o == nil {
		return nil
	}
	return o.AccessToken
}

func (o *PropertyConnectionAuth) GetAPIURL() *string {
	if o == nil {
		return nil
	}
	return o.APIURL
}

func (o *PropertyConnectionAuth) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *PropertyConnectionAuth) GetAuthorizeURL() *string {
	if o == nil {
		return nil
	}
	return o.AuthorizeURL
}

func (o *PropertyConnectionAuth) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *PropertyConnectionAuth) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}

func (o *PropertyConnectionAuth) GetConsumerKey() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerKey
}

func (o *PropertyConnectionAuth) GetConsumerSecret() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerSecret
}

func (o *PropertyConnectionAuth) GetDevAPIKey() *string {
	if o == nil {
		return nil
	}
	return o.DevAPIKey
}

func (o *PropertyConnectionAuth) GetEmails() []string {
	if o == nil {
		return nil
	}
	return o.Emails
}

func (o *PropertyConnectionAuth) GetExpiresIn() *float64 {
	if o == nil {
		return nil
	}
	return o.ExpiresIn
}

func (o *PropertyConnectionAuth) GetExpiryDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.ExpiryDate
}

func (o *PropertyConnectionAuth) GetKey() *string {
	if o == nil {
		return nil
	}
	return o.Key
}

func (o *PropertyConnectionAuth) GetMeta() map[string]any {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *PropertyConnectionAuth) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PropertyConnectionAuth) GetOtherAuthInfo() []string {
	if o == nil {
		return nil
	}
	return o.OtherAuthInfo
}

func (o *PropertyConnectionAuth) GetPem() *string {
	if o == nil {
		return nil
	}
	return o.Pem
}

func (o *PropertyConnectionAuth) GetRefreshToken() *string {
	if o == nil {
		return nil
	}
	return o.RefreshToken
}

func (o *PropertyConnectionAuth) GetRefreshTokenExpiresDate() *time.Time {
	if o == nil {
		return nil
	}
	return o.RefreshTokenExpiresDate
}

func (o *PropertyConnectionAuth) GetRefreshTokenExpiresIn() *float64 {
	if o == nil {
		return nil
	}
	return o.RefreshTokenExpiresIn
}

func (o *PropertyConnectionAuth) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}

func (o *PropertyConnectionAuth) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

func (o *PropertyConnectionAuth) GetTokenURL() *string {
	if o == nil {
		return nil
	}
	return o.TokenURL
}

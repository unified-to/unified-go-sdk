// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

// MarketingList - Mailing List
type MarketingList struct {
	CreatedAt *string `json:"created_at,omitempty"`
	ID        *string `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
	// The raw data returned by the integration for this list
	Raw       *PropertyMarketingListRaw `json:"raw,omitempty"`
	UpdatedAt *types.Date               `json:"updated_at,omitempty"`
}

func (o *MarketingList) GetCreatedAt() *string {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *MarketingList) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *MarketingList) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *MarketingList) GetRaw() *PropertyMarketingListRaw {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *MarketingList) GetUpdatedAt() *types.Date {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

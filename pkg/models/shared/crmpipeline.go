// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/types"
)

type CrmPipeline struct {
	Active          *bool                   `json:"active,omitempty"`
	CreatedAt       *types.Date             `json:"created_at,omitempty"`
	DealProbability *bool                   `json:"deal_probability,omitempty"`
	DisplayOrder    *float64                `json:"display_order,omitempty"`
	ID              *string                 `json:"id,omitempty"`
	Name            *string                 `json:"name,omitempty"`
	Raw             *PropertyCrmPipelineRaw `json:"raw,omitempty"`
	UpdatedAt       *types.Date             `json:"updated_at,omitempty"`
}

func (o *CrmPipeline) GetActive() *bool {
	if o == nil {
		return nil
	}
	return o.Active
}

func (o *CrmPipeline) GetCreatedAt() *types.Date {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CrmPipeline) GetDealProbability() *bool {
	if o == nil {
		return nil
	}
	return o.DealProbability
}

func (o *CrmPipeline) GetDisplayOrder() *float64 {
	if o == nil {
		return nil
	}
	return o.DisplayOrder
}

func (o *CrmPipeline) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CrmPipeline) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CrmPipeline) GetRaw() *PropertyCrmPipelineRaw {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *CrmPipeline) GetUpdatedAt() *types.Date {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

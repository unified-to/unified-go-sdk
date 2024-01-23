// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type CrmPipeline struct {
	CreatedAt       *time.Time              `json:"created_at,omitempty"`
	DealProbability *float64                `json:"deal_probability,omitempty"`
	DisplayOrder    *float64                `json:"display_order,omitempty"`
	ID              *string                 `json:"id,omitempty"`
	IsActive        *bool                   `json:"is_active,omitempty"`
	Name            *string                 `json:"name,omitempty"`
	Raw             *PropertyCrmPipelineRaw `json:"raw,omitempty"`
	UpdatedAt       *time.Time              `json:"updated_at,omitempty"`
}

func (c CrmPipeline) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CrmPipeline) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CrmPipeline) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CrmPipeline) GetDealProbability() *float64 {
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

func (o *CrmPipeline) GetIsActive() *bool {
	if o == nil {
		return nil
	}
	return o.IsActive
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

func (o *CrmPipeline) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

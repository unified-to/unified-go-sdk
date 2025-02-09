// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type CrmDealRaw struct {
}

// CrmDeal - A deal represents an opportunity with companies and/or contacts
type CrmDeal struct {
	Amount      *float64    `json:"amount,omitempty"`
	ClosedAt    *time.Time  `json:"closed_at,omitempty"`
	CompanyIds  []string    `json:"company_ids,omitempty"`
	ContactIds  []string    `json:"contact_ids,omitempty"`
	CreatedAt   *time.Time  `json:"created_at,omitempty"`
	Currency    *string     `json:"currency,omitempty"`
	ID          *string     `json:"id,omitempty"`
	LostReason  *string     `json:"lost_reason,omitempty"`
	Name        *string     `json:"name,omitempty"`
	Pipeline    *string     `json:"pipeline,omitempty"`
	PipelineID  *string     `json:"pipeline_id,omitempty"`
	Probability *float64    `json:"probability,omitempty"`
	Raw         *CrmDealRaw `json:"raw,omitempty"`
	Source      *string     `json:"source,omitempty"`
	Stage       *string     `json:"stage,omitempty"`
	StageID     *string     `json:"stage_id,omitempty"`
	Tags        []string    `json:"tags,omitempty"`
	UpdatedAt   *time.Time  `json:"updated_at,omitempty"`
	UserID      *string     `json:"user_id,omitempty"`
	WonReason   *string     `json:"won_reason,omitempty"`
}

func (c CrmDeal) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CrmDeal) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CrmDeal) GetAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *CrmDeal) GetClosedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ClosedAt
}

func (o *CrmDeal) GetCompanyIds() []string {
	if o == nil {
		return nil
	}
	return o.CompanyIds
}

func (o *CrmDeal) GetContactIds() []string {
	if o == nil {
		return nil
	}
	return o.ContactIds
}

func (o *CrmDeal) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *CrmDeal) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *CrmDeal) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *CrmDeal) GetLostReason() *string {
	if o == nil {
		return nil
	}
	return o.LostReason
}

func (o *CrmDeal) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *CrmDeal) GetPipeline() *string {
	if o == nil {
		return nil
	}
	return o.Pipeline
}

func (o *CrmDeal) GetPipelineID() *string {
	if o == nil {
		return nil
	}
	return o.PipelineID
}

func (o *CrmDeal) GetProbability() *float64 {
	if o == nil {
		return nil
	}
	return o.Probability
}

func (o *CrmDeal) GetRaw() *CrmDealRaw {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *CrmDeal) GetSource() *string {
	if o == nil {
		return nil
	}
	return o.Source
}

func (o *CrmDeal) GetStage() *string {
	if o == nil {
		return nil
	}
	return o.Stage
}

func (o *CrmDeal) GetStageID() *string {
	if o == nil {
		return nil
	}
	return o.StageID
}

func (o *CrmDeal) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CrmDeal) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *CrmDeal) GetUserID() *string {
	if o == nil {
		return nil
	}
	return o.UserID
}

func (o *CrmDeal) GetWonReason() *string {
	if o == nil {
		return nil
	}
	return o.WonReason
}

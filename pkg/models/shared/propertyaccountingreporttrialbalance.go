// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type PropertyAccountingReportTrialBalance struct {
	CreatedAt         *time.Time                      `json:"created_at,omitempty"`
	Currency          *string                         `json:"currency,omitempty"`
	EndAt             *time.Time                      `json:"end_at,omitempty"`
	ID                *string                         `json:"id,omitempty"`
	Name              *string                         `json:"name,omitempty"`
	Raw               map[string]any                  `json:"raw,omitempty"`
	StartAt           *time.Time                      `json:"start_at,omitempty"`
	SubItems          []AccountingTrialbalanceSubItem `json:"sub_items,omitempty"`
	TotalCreditAmount *float64                        `json:"total_credit_amount,omitempty"`
	TotalDebitAmount  *float64                        `json:"total_debit_amount,omitempty"`
	UpdatedAt         *time.Time                      `json:"updated_at,omitempty"`
}

func (p PropertyAccountingReportTrialBalance) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PropertyAccountingReportTrialBalance) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PropertyAccountingReportTrialBalance) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *PropertyAccountingReportTrialBalance) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *PropertyAccountingReportTrialBalance) GetEndAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndAt
}

func (o *PropertyAccountingReportTrialBalance) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PropertyAccountingReportTrialBalance) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PropertyAccountingReportTrialBalance) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *PropertyAccountingReportTrialBalance) GetStartAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartAt
}

func (o *PropertyAccountingReportTrialBalance) GetSubItems() []AccountingTrialbalanceSubItem {
	if o == nil {
		return nil
	}
	return o.SubItems
}

func (o *PropertyAccountingReportTrialBalance) GetTotalCreditAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalCreditAmount
}

func (o *PropertyAccountingReportTrialBalance) GetTotalDebitAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.TotalDebitAmount
}

func (o *PropertyAccountingReportTrialBalance) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

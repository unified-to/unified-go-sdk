// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type PropertyAccountingReportBalanceSheet struct {
	Assets          []AccountingBalancesheetItem `json:"assets,omitempty"`
	CreatedAt       *time.Time                   `json:"created_at,omitempty"`
	Currency        *string                      `json:"currency,omitempty"`
	EndAt           *time.Time                   `json:"end_at,omitempty"`
	Equity          []AccountingBalancesheetItem `json:"equity,omitempty"`
	ID              *string                      `json:"id,omitempty"`
	Liabilities     []AccountingBalancesheetItem `json:"liabilities,omitempty"`
	Name            *string                      `json:"name,omitempty"`
	NetAssetsAmount *float64                     `json:"net_assets_amount,omitempty"`
	Raw             map[string]any               `json:"raw,omitempty"`
	StartAt         *time.Time                   `json:"start_at,omitempty"`
	UpdatedAt       *time.Time                   `json:"updated_at,omitempty"`
}

func (p PropertyAccountingReportBalanceSheet) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(p, "", false)
}

func (p *PropertyAccountingReportBalanceSheet) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &p, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *PropertyAccountingReportBalanceSheet) GetAssets() []AccountingBalancesheetItem {
	if o == nil {
		return nil
	}
	return o.Assets
}

func (o *PropertyAccountingReportBalanceSheet) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *PropertyAccountingReportBalanceSheet) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *PropertyAccountingReportBalanceSheet) GetEndAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.EndAt
}

func (o *PropertyAccountingReportBalanceSheet) GetEquity() []AccountingBalancesheetItem {
	if o == nil {
		return nil
	}
	return o.Equity
}

func (o *PropertyAccountingReportBalanceSheet) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *PropertyAccountingReportBalanceSheet) GetLiabilities() []AccountingBalancesheetItem {
	if o == nil {
		return nil
	}
	return o.Liabilities
}

func (o *PropertyAccountingReportBalanceSheet) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PropertyAccountingReportBalanceSheet) GetNetAssetsAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.NetAssetsAmount
}

func (o *PropertyAccountingReportBalanceSheet) GetRaw() map[string]any {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *PropertyAccountingReportBalanceSheet) GetStartAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.StartAt
}

func (o *PropertyAccountingReportBalanceSheet) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

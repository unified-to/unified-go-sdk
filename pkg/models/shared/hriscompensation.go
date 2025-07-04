// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type HrisCompensationFrequency string

const (
	HrisCompensationFrequencyOneTime HrisCompensationFrequency = "ONE_TIME"
	HrisCompensationFrequencyDay     HrisCompensationFrequency = "DAY"
	HrisCompensationFrequencyQuarter HrisCompensationFrequency = "QUARTER"
	HrisCompensationFrequencyYear    HrisCompensationFrequency = "YEAR"
	HrisCompensationFrequencyHour    HrisCompensationFrequency = "HOUR"
	HrisCompensationFrequencyMonth   HrisCompensationFrequency = "MONTH"
	HrisCompensationFrequencyWeek    HrisCompensationFrequency = "WEEK"
)

func (e HrisCompensationFrequency) ToPointer() *HrisCompensationFrequency {
	return &e
}

type HrisCompensationType string

const (
	HrisCompensationTypeSalary       HrisCompensationType = "SALARY"
	HrisCompensationTypeBonus        HrisCompensationType = "BONUS"
	HrisCompensationTypeStockOptions HrisCompensationType = "STOCK_OPTIONS"
	HrisCompensationTypeEquity       HrisCompensationType = "EQUITY"
	HrisCompensationTypeOther        HrisCompensationType = "OTHER"
)

func (e HrisCompensationType) ToPointer() *HrisCompensationType {
	return &e
}

type HrisCompensation struct {
	Amount    *float64                   `json:"amount,omitempty"`
	Currency  *string                    `json:"currency,omitempty"`
	Frequency *HrisCompensationFrequency `json:"frequency,omitempty"`
	Type      *HrisCompensationType      `json:"type,omitempty"`
}

func (o *HrisCompensation) GetAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *HrisCompensation) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *HrisCompensation) GetFrequency() *HrisCompensationFrequency {
	if o == nil {
		return nil
	}
	return o.Frequency
}

func (o *HrisCompensation) GetType() *HrisCompensationType {
	if o == nil {
		return nil
	}
	return o.Type
}

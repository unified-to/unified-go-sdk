// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type Frequency string

const (
	FrequencyOneTime Frequency = "ONE_TIME"
	FrequencyDay     Frequency = "DAY"
	FrequencyQuarter Frequency = "QUARTER"
	FrequencyYear    Frequency = "YEAR"
	FrequencyHour    Frequency = "HOUR"
	FrequencyMonth   Frequency = "MONTH"
	FrequencyWeek    Frequency = "WEEK"
)

func (e Frequency) ToPointer() *Frequency {
	return &e
}

type AtsCompensationType string

const (
	AtsCompensationTypeSalary       AtsCompensationType = "SALARY"
	AtsCompensationTypeBonus        AtsCompensationType = "BONUS"
	AtsCompensationTypeStockOptions AtsCompensationType = "STOCK_OPTIONS"
	AtsCompensationTypeEquity       AtsCompensationType = "EQUITY"
	AtsCompensationTypeOther        AtsCompensationType = "OTHER"
)

func (e AtsCompensationType) ToPointer() *AtsCompensationType {
	return &e
}

type AtsCompensation struct {
	Currency  *string              `json:"currency,omitempty"`
	Frequency *Frequency           `json:"frequency,omitempty"`
	Max       *float64             `json:"max,omitempty"`
	Min       *float64             `json:"min,omitempty"`
	Type      *AtsCompensationType `json:"type,omitempty"`
}

func (o *AtsCompensation) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AtsCompensation) GetFrequency() *Frequency {
	if o == nil {
		return nil
	}
	return o.Frequency
}

func (o *AtsCompensation) GetMax() *float64 {
	if o == nil {
		return nil
	}
	return o.Max
}

func (o *AtsCompensation) GetMin() *float64 {
	if o == nil {
		return nil
	}
	return o.Min
}

func (o *AtsCompensation) GetType() *AtsCompensationType {
	if o == nil {
		return nil
	}
	return o.Type
}

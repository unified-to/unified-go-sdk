// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

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
func (e *HrisCompensationFrequency) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ONE_TIME":
		fallthrough
	case "DAY":
		fallthrough
	case "QUARTER":
		fallthrough
	case "YEAR":
		fallthrough
	case "HOUR":
		fallthrough
	case "MONTH":
		fallthrough
	case "WEEK":
		*e = HrisCompensationFrequency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HrisCompensationFrequency: %v", v)
	}
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
func (e *HrisCompensationType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "SALARY":
		fallthrough
	case "BONUS":
		fallthrough
	case "STOCK_OPTIONS":
		fallthrough
	case "EQUITY":
		fallthrough
	case "OTHER":
		*e = HrisCompensationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HrisCompensationType: %v", v)
	}
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

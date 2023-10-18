// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type AtsCompensationFrequency string

const (
	AtsCompensationFrequencyOneTime AtsCompensationFrequency = "ONE_TIME"
	AtsCompensationFrequencyDay     AtsCompensationFrequency = "DAY"
	AtsCompensationFrequencyQuarter AtsCompensationFrequency = "QUARTER"
	AtsCompensationFrequencyYear    AtsCompensationFrequency = "YEAR"
	AtsCompensationFrequencyHour    AtsCompensationFrequency = "HOUR"
	AtsCompensationFrequencyMonth   AtsCompensationFrequency = "MONTH"
	AtsCompensationFrequencyWeek    AtsCompensationFrequency = "WEEK"
)

func (e AtsCompensationFrequency) ToPointer() *AtsCompensationFrequency {
	return &e
}

func (e *AtsCompensationFrequency) UnmarshalJSON(data []byte) error {
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
		*e = AtsCompensationFrequency(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AtsCompensationFrequency: %v", v)
	}
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

func (e *AtsCompensationType) UnmarshalJSON(data []byte) error {
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
		*e = AtsCompensationType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AtsCompensationType: %v", v)
	}
}

type AtsCompensation struct {
	Currency  *string                   `json:"currency,omitempty"`
	Frequency *AtsCompensationFrequency `json:"frequency,omitempty"`
	Max       *float64                  `json:"max,omitempty"`
	Min       *float64                  `json:"min,omitempty"`
	Type      AtsCompensationType       `json:"type"`
}

func (o *AtsCompensation) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *AtsCompensation) GetFrequency() *AtsCompensationFrequency {
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

func (o *AtsCompensation) GetType() AtsCompensationType {
	if o == nil {
		return AtsCompensationType("")
	}
	return o.Type
}

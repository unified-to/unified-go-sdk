// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ScimEmailType string

const (
	ScimEmailTypeWork  ScimEmailType = "work"
	ScimEmailTypeHome  ScimEmailType = "home"
	ScimEmailTypeOther ScimEmailType = "other"
)

func (e ScimEmailType) ToPointer() *ScimEmailType {
	return &e
}
func (e *ScimEmailType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "work":
		fallthrough
	case "home":
		fallthrough
	case "other":
		*e = ScimEmailType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ScimEmailType: %v", v)
	}
}

type ScimEmail struct {
	Display *string       `json:"display,omitempty"`
	Primary *bool         `json:"primary,omitempty"`
	Type    ScimEmailType `json:"type"`
	Value   *string       `json:"value,omitempty"`
}

func (o *ScimEmail) GetDisplay() *string {
	if o == nil {
		return nil
	}
	return o.Display
}

func (o *ScimEmail) GetPrimary() *bool {
	if o == nil {
		return nil
	}
	return o.Primary
}

func (o *ScimEmail) GetType() ScimEmailType {
	if o == nil {
		return ScimEmailType("")
	}
	return o.Type
}

func (o *ScimEmail) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}
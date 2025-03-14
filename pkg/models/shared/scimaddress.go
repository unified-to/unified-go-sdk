// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ScimAddressType string

const (
	ScimAddressTypeWork  ScimAddressType = "work"
	ScimAddressTypeHome  ScimAddressType = "home"
	ScimAddressTypeOther ScimAddressType = "other"
)

func (e ScimAddressType) ToPointer() *ScimAddressType {
	return &e
}
func (e *ScimAddressType) UnmarshalJSON(data []byte) error {
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
		*e = ScimAddressType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ScimAddressType: %v", v)
	}
}

type ScimAddress struct {
	Country       *string          `json:"country,omitempty"`
	Formatted     *string          `json:"formatted,omitempty"`
	Locality      *string          `json:"locality,omitempty"`
	PostalCode    *string          `json:"postalCode,omitempty"`
	Region        *string          `json:"region,omitempty"`
	StreetAddress *string          `json:"streetAddress,omitempty"`
	Type          *ScimAddressType `json:"type,omitempty"`
}

func (o *ScimAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *ScimAddress) GetFormatted() *string {
	if o == nil {
		return nil
	}
	return o.Formatted
}

func (o *ScimAddress) GetLocality() *string {
	if o == nil {
		return nil
	}
	return o.Locality
}

func (o *ScimAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *ScimAddress) GetRegion() *string {
	if o == nil {
		return nil
	}
	return o.Region
}

func (o *ScimAddress) GetStreetAddress() *string {
	if o == nil {
		return nil
	}
	return o.StreetAddress
}

func (o *ScimAddress) GetType() *ScimAddressType {
	if o == nil {
		return nil
	}
	return o.Type
}

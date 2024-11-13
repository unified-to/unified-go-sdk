// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ScimImsType string

const (
	ScimImsTypeAim   ScimImsType = "aim"
	ScimImsTypeQtalk ScimImsType = "qtalk"
	ScimImsTypeIcq   ScimImsType = "icq"
	ScimImsTypeXMPP  ScimImsType = "xmpp"
	ScimImsTypeMsn   ScimImsType = "msn"
	ScimImsTypeSkype ScimImsType = "skype"
	ScimImsTypeQq    ScimImsType = "qq"
	ScimImsTypeYahoo ScimImsType = "yahoo"
)

func (e ScimImsType) ToPointer() *ScimImsType {
	return &e
}
func (e *ScimImsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "aim":
		fallthrough
	case "qtalk":
		fallthrough
	case "icq":
		fallthrough
	case "xmpp":
		fallthrough
	case "msn":
		fallthrough
	case "skype":
		fallthrough
	case "qq":
		fallthrough
	case "yahoo":
		*e = ScimImsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ScimImsType: %v", v)
	}
}

type ScimIms struct {
	Display *string      `json:"display,omitempty"`
	Primary *bool        `json:"primary,omitempty"`
	Type    *ScimImsType `json:"type,omitempty"`
	Value   *string      `json:"value,omitempty"`
}

func (o *ScimIms) GetDisplay() *string {
	if o == nil {
		return nil
	}
	return o.Display
}

func (o *ScimIms) GetPrimary() *bool {
	if o == nil {
		return nil
	}
	return o.Primary
}

func (o *ScimIms) GetType() *ScimImsType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ScimIms) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}
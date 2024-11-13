// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ScimUserGroupsType string

const (
	ScimUserGroupsTypeDirect   ScimUserGroupsType = "direct"
	ScimUserGroupsTypeIndirect ScimUserGroupsType = "indirect"
)

func (e ScimUserGroupsType) ToPointer() *ScimUserGroupsType {
	return &e
}
func (e *ScimUserGroupsType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "direct":
		fallthrough
	case "indirect":
		*e = ScimUserGroupsType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ScimUserGroupsType: %v", v)
	}
}

type ScimUserGroups struct {
	DollarRef *string             `json:"$ref,omitempty"`
	Display   *string             `json:"display,omitempty"`
	Type      *ScimUserGroupsType `json:"type,omitempty"`
	Value     string              `json:"value"`
}

func (o *ScimUserGroups) GetDollarRef() *string {
	if o == nil {
		return nil
	}
	return o.DollarRef
}

func (o *ScimUserGroups) GetDisplay() *string {
	if o == nil {
		return nil
	}
	return o.Display
}

func (o *ScimUserGroups) GetType() *ScimUserGroupsType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ScimUserGroups) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}
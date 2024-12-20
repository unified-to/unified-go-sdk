// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Operation string

const (
	OperationAdd    Operation = "add"
	OperationDelete Operation = "delete"
)

func (e Operation) ToPointer() *Operation {
	return &e
}
func (e *Operation) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "add":
		fallthrough
	case "delete":
		*e = Operation(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Operation: %v", v)
	}
}

type ScimGroupMemberType string

const (
	ScimGroupMemberTypeUser  ScimGroupMemberType = "User"
	ScimGroupMemberTypeGroup ScimGroupMemberType = "Group"
)

func (e ScimGroupMemberType) ToPointer() *ScimGroupMemberType {
	return &e
}
func (e *ScimGroupMemberType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "User":
		fallthrough
	case "Group":
		*e = ScimGroupMemberType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ScimGroupMemberType: %v", v)
	}
}

type ScimGroupMember struct {
	DollarRef *string              `json:"$ref,omitempty"`
	Display   *string              `json:"display,omitempty"`
	Operation *Operation           `json:"operation,omitempty"`
	Type      *ScimGroupMemberType `json:"type,omitempty"`
	Value     string               `json:"value"`
}

func (o *ScimGroupMember) GetDollarRef() *string {
	if o == nil {
		return nil
	}
	return o.DollarRef
}

func (o *ScimGroupMember) GetDisplay() *string {
	if o == nil {
		return nil
	}
	return o.Display
}

func (o *ScimGroupMember) GetOperation() *Operation {
	if o == nil {
		return nil
	}
	return o.Operation
}

func (o *ScimGroupMember) GetType() *ScimGroupMemberType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ScimGroupMember) GetValue() string {
	if o == nil {
		return ""
	}
	return o.Value
}

// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PropertyPropertyUserUrnIetfParamsScimSchemasExtensionEnterprise20UserManagerType string

const (
	PropertyPropertyUserUrnIetfParamsScimSchemasExtensionEnterprise20UserManagerTypeDirect   PropertyPropertyUserUrnIetfParamsScimSchemasExtensionEnterprise20UserManagerType = "direct"
	PropertyPropertyUserUrnIetfParamsScimSchemasExtensionEnterprise20UserManagerTypeIndirect PropertyPropertyUserUrnIetfParamsScimSchemasExtensionEnterprise20UserManagerType = "indirect"
)

func (e PropertyPropertyUserUrnIetfParamsScimSchemasExtensionEnterprise20UserManagerType) ToPointer() *PropertyPropertyUserUrnIetfParamsScimSchemasExtensionEnterprise20UserManagerType {
	return &e
}
func (e *PropertyPropertyUserUrnIetfParamsScimSchemasExtensionEnterprise20UserManagerType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "direct":
		fallthrough
	case "indirect":
		*e = PropertyPropertyUserUrnIetfParamsScimSchemasExtensionEnterprise20UserManagerType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PropertyPropertyUserUrnIetfParamsScimSchemasExtensionEnterprise20UserManagerType: %v", v)
	}
}

type PropertyPropertyUserUrnIetfParamsScimSchemasExtensionEnterprise20UserManager struct {
	DollarRef   *string                                                                           `json:"$ref,omitempty"`
	DisplayName *string                                                                           `json:"displayName,omitempty"`
	ManagerID   *string                                                                           `json:"managerId,omitempty"`
	Type        *PropertyPropertyUserUrnIetfParamsScimSchemasExtensionEnterprise20UserManagerType `json:"type,omitempty"`
	Value       *string                                                                           `json:"value,omitempty"`
}

func (o *PropertyPropertyUserUrnIetfParamsScimSchemasExtensionEnterprise20UserManager) GetDollarRef() *string {
	if o == nil {
		return nil
	}
	return o.DollarRef
}

func (o *PropertyPropertyUserUrnIetfParamsScimSchemasExtensionEnterprise20UserManager) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *PropertyPropertyUserUrnIetfParamsScimSchemasExtensionEnterprise20UserManager) GetManagerID() *string {
	if o == nil {
		return nil
	}
	return o.ManagerID
}

func (o *PropertyPropertyUserUrnIetfParamsScimSchemasExtensionEnterprise20UserManager) GetType() *PropertyPropertyUserUrnIetfParamsScimSchemasExtensionEnterprise20UserManagerType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *PropertyPropertyUserUrnIetfParamsScimSchemasExtensionEnterprise20UserManager) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}
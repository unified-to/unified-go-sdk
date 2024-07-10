// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PropertyHrisEmployeeEmployeeRoles string

const (
	PropertyHrisEmployeeEmployeeRolesAdmin       PropertyHrisEmployeeEmployeeRoles = "ADMIN"
	PropertyHrisEmployeeEmployeeRolesManager     PropertyHrisEmployeeEmployeeRoles = "MANAGER"
	PropertyHrisEmployeeEmployeeRolesRecruiter   PropertyHrisEmployeeEmployeeRoles = "RECRUITER"
	PropertyHrisEmployeeEmployeeRolesSalesrep    PropertyHrisEmployeeEmployeeRoles = "SALESREP"
	PropertyHrisEmployeeEmployeeRolesInterviewer PropertyHrisEmployeeEmployeeRoles = "INTERVIEWER"
)

func (e PropertyHrisEmployeeEmployeeRoles) ToPointer() *PropertyHrisEmployeeEmployeeRoles {
	return &e
}
func (e *PropertyHrisEmployeeEmployeeRoles) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ADMIN":
		fallthrough
	case "MANAGER":
		fallthrough
	case "RECRUITER":
		fallthrough
	case "SALESREP":
		fallthrough
	case "INTERVIEWER":
		*e = PropertyHrisEmployeeEmployeeRoles(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PropertyHrisEmployeeEmployeeRoles: %v", v)
	}
}
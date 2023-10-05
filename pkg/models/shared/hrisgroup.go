// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type HrisGroupType string

const (
	HrisGroupTypeTeam          HrisGroupType = "TEAM"
	HrisGroupTypeGroup         HrisGroupType = "GROUP"
	HrisGroupTypeDepartment    HrisGroupType = "DEPARTMENT"
	HrisGroupTypeDivision      HrisGroupType = "DIVISION"
	HrisGroupTypeBusinessUnit  HrisGroupType = "BUSINESS_UNIT"
	HrisGroupTypeBranch        HrisGroupType = "BRANCH"
	HrisGroupTypeSubDepartment HrisGroupType = "SUB_DEPARTMENT"
)

func (e HrisGroupType) ToPointer() *HrisGroupType {
	return &e
}

func (e *HrisGroupType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "TEAM":
		fallthrough
	case "GROUP":
		fallthrough
	case "DEPARTMENT":
		fallthrough
	case "DIVISION":
		fallthrough
	case "BUSINESS_UNIT":
		fallthrough
	case "BRANCH":
		fallthrough
	case "SUB_DEPARTMENT":
		*e = HrisGroupType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HrisGroupType: %v", v)
	}
}

type HrisGroup struct {
	CreatedAt   *time.Time            `json:"created_at,omitempty"`
	Description *string               `json:"description,omitempty"`
	EmployeeIds []string              `json:"employee_ids,omitempty"`
	ID          *string               `json:"id,omitempty"`
	IsActive    *bool                 `json:"is_active,omitempty"`
	ManagerIds  []string              `json:"manager_ids,omitempty"`
	Name        *string               `json:"name,omitempty"`
	ParentID    *string               `json:"parent_id,omitempty"`
	Raw         *PropertyHrisGroupRaw `json:"raw,omitempty"`
	Type        *HrisGroupType        `json:"type,omitempty"`
	UpdatedAt   *time.Time            `json:"updated_at,omitempty"`
}

func (h HrisGroup) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HrisGroup) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *HrisGroup) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *HrisGroup) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *HrisGroup) GetEmployeeIds() []string {
	if o == nil {
		return nil
	}
	return o.EmployeeIds
}

func (o *HrisGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *HrisGroup) GetIsActive() *bool {
	if o == nil {
		return nil
	}
	return o.IsActive
}

func (o *HrisGroup) GetManagerIds() []string {
	if o == nil {
		return nil
	}
	return o.ManagerIds
}

func (o *HrisGroup) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *HrisGroup) GetParentID() *string {
	if o == nil {
		return nil
	}
	return o.ParentID
}

func (o *HrisGroup) GetRaw() *PropertyHrisGroupRaw {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *HrisGroup) GetType() *HrisGroupType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *HrisGroup) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}
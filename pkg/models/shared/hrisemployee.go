// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/unified-to/unified-go-sdk/pkg/utils"
	"time"
)

type EmploymentStatus string

const (
	EmploymentStatusActive   EmploymentStatus = "ACTIVE"
	EmploymentStatusInactive EmploymentStatus = "INACTIVE"
)

func (e EmploymentStatus) ToPointer() *EmploymentStatus {
	return &e
}
func (e *EmploymentStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ACTIVE":
		fallthrough
	case "INACTIVE":
		*e = EmploymentStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for EmploymentStatus: %v", v)
	}
}

type HrisEmployeeEmploymentType string

const (
	HrisEmployeeEmploymentTypeFullTime   HrisEmployeeEmploymentType = "FULL_TIME"
	HrisEmployeeEmploymentTypePartTime   HrisEmployeeEmploymentType = "PART_TIME"
	HrisEmployeeEmploymentTypeContractor HrisEmployeeEmploymentType = "CONTRACTOR"
	HrisEmployeeEmploymentTypeIntern     HrisEmployeeEmploymentType = "INTERN"
	HrisEmployeeEmploymentTypeConsultant HrisEmployeeEmploymentType = "CONSULTANT"
	HrisEmployeeEmploymentTypeVolunteer  HrisEmployeeEmploymentType = "VOLUNTEER"
	HrisEmployeeEmploymentTypeCasual     HrisEmployeeEmploymentType = "CASUAL"
	HrisEmployeeEmploymentTypeSeasonal   HrisEmployeeEmploymentType = "SEASONAL"
	HrisEmployeeEmploymentTypeFreelance  HrisEmployeeEmploymentType = "FREELANCE"
	HrisEmployeeEmploymentTypeOther      HrisEmployeeEmploymentType = "OTHER"
)

func (e HrisEmployeeEmploymentType) ToPointer() *HrisEmployeeEmploymentType {
	return &e
}
func (e *HrisEmployeeEmploymentType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "FULL_TIME":
		fallthrough
	case "PART_TIME":
		fallthrough
	case "CONTRACTOR":
		fallthrough
	case "INTERN":
		fallthrough
	case "CONSULTANT":
		fallthrough
	case "VOLUNTEER":
		fallthrough
	case "CASUAL":
		fallthrough
	case "SEASONAL":
		fallthrough
	case "FREELANCE":
		fallthrough
	case "OTHER":
		*e = HrisEmployeeEmploymentType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HrisEmployeeEmploymentType: %v", v)
	}
}

type HrisEmployeeGender string

const (
	HrisEmployeeGenderMale      HrisEmployeeGender = "MALE"
	HrisEmployeeGenderFemale    HrisEmployeeGender = "FEMALE"
	HrisEmployeeGenderIntersex  HrisEmployeeGender = "INTERSEX"
	HrisEmployeeGenderTrans     HrisEmployeeGender = "TRANS"
	HrisEmployeeGenderNonBinary HrisEmployeeGender = "NON_BINARY"
)

func (e HrisEmployeeGender) ToPointer() *HrisEmployeeGender {
	return &e
}
func (e *HrisEmployeeGender) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MALE":
		fallthrough
	case "FEMALE":
		fallthrough
	case "INTERSEX":
		fallthrough
	case "TRANS":
		fallthrough
	case "NON_BINARY":
		*e = HrisEmployeeGender(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HrisEmployeeGender: %v", v)
	}
}

type MaritalStatus string

const (
	MaritalStatusMarried MaritalStatus = "MARRIED"
	MaritalStatusSingle  MaritalStatus = "SINGLE"
)

func (e MaritalStatus) ToPointer() *MaritalStatus {
	return &e
}
func (e *MaritalStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "MARRIED":
		fallthrough
	case "SINGLE":
		*e = MaritalStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MaritalStatus: %v", v)
	}
}

type HrisEmployeeRaw struct {
}

type HrisEmployee struct {
	Address          *PropertyHrisEmployeeAddress        `json:"address,omitempty"`
	Bio              *string                             `json:"bio,omitempty"`
	CompanyID        *string                             `json:"company_id,omitempty"`
	Compensation     []HrisCompensation                  `json:"compensation,omitempty"`
	CreatedAt        *time.Time                          `json:"created_at,omitempty"`
	Currency         *string                             `json:"currency,omitempty"`
	DateOfBirth      *time.Time                          `json:"date_of_birth,omitempty"`
	Department       *string                             `json:"department,omitempty"`
	Division         *string                             `json:"division,omitempty"`
	Emails           []HrisEmail                         `json:"emails,omitempty"`
	EmployeeNumber   *string                             `json:"employee_number,omitempty"`
	EmployeeRoles    []PropertyHrisEmployeeEmployeeRoles `json:"employee_roles,omitempty"`
	EmploymentStatus *EmploymentStatus                   `json:"employment_status,omitempty"`
	EmploymentType   *HrisEmployeeEmploymentType         `json:"employment_type,omitempty"`
	Gender           *HrisEmployeeGender                 `json:"gender,omitempty"`
	// Which groups/teams/units that this employee/user belongs to.  May not have all of the Group fields present, but should have id, name, or email.
	Groups         []HrisGroup      `json:"groups,omitempty"`
	HiredAt        *time.Time       `json:"hired_at,omitempty"`
	ID             *string          `json:"id,omitempty"`
	ImageURL       *string          `json:"image_url,omitempty"`
	LanguageLocale *string          `json:"language_locale,omitempty"`
	Location       *string          `json:"location,omitempty"`
	Locations      []HrisLocation   `json:"locations,omitempty"`
	ManagerID      *string          `json:"manager_id,omitempty"`
	MaritalStatus  *MaritalStatus   `json:"marital_status,omitempty"`
	Metadata       []HrisMetadata   `json:"metadata,omitempty"`
	Name           *string          `json:"name,omitempty"`
	Pronouns       *string          `json:"pronouns,omitempty"`
	Raw            *HrisEmployeeRaw `json:"raw,omitempty"`
	Salutation     *string          `json:"salutation,omitempty"`
	SsnSin         *string          `json:"ssn_sin,omitempty"`
	Telephones     []HrisTelephone  `json:"telephones,omitempty"`
	TerminatedAt   *time.Time       `json:"terminated_at,omitempty"`
	Timezone       *string          `json:"timezone,omitempty"`
	Title          *string          `json:"title,omitempty"`
	UpdatedAt      *time.Time       `json:"updated_at,omitempty"`
}

func (h HrisEmployee) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(h, "", false)
}

func (h *HrisEmployee) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &h, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *HrisEmployee) GetAddress() *PropertyHrisEmployeeAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *HrisEmployee) GetBio() *string {
	if o == nil {
		return nil
	}
	return o.Bio
}

func (o *HrisEmployee) GetCompanyID() *string {
	if o == nil {
		return nil
	}
	return o.CompanyID
}

func (o *HrisEmployee) GetCompensation() []HrisCompensation {
	if o == nil {
		return nil
	}
	return o.Compensation
}

func (o *HrisEmployee) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *HrisEmployee) GetCurrency() *string {
	if o == nil {
		return nil
	}
	return o.Currency
}

func (o *HrisEmployee) GetDateOfBirth() *time.Time {
	if o == nil {
		return nil
	}
	return o.DateOfBirth
}

func (o *HrisEmployee) GetDepartment() *string {
	if o == nil {
		return nil
	}
	return o.Department
}

func (o *HrisEmployee) GetDivision() *string {
	if o == nil {
		return nil
	}
	return o.Division
}

func (o *HrisEmployee) GetEmails() []HrisEmail {
	if o == nil {
		return nil
	}
	return o.Emails
}

func (o *HrisEmployee) GetEmployeeNumber() *string {
	if o == nil {
		return nil
	}
	return o.EmployeeNumber
}

func (o *HrisEmployee) GetEmployeeRoles() []PropertyHrisEmployeeEmployeeRoles {
	if o == nil {
		return nil
	}
	return o.EmployeeRoles
}

func (o *HrisEmployee) GetEmploymentStatus() *EmploymentStatus {
	if o == nil {
		return nil
	}
	return o.EmploymentStatus
}

func (o *HrisEmployee) GetEmploymentType() *HrisEmployeeEmploymentType {
	if o == nil {
		return nil
	}
	return o.EmploymentType
}

func (o *HrisEmployee) GetGender() *HrisEmployeeGender {
	if o == nil {
		return nil
	}
	return o.Gender
}

func (o *HrisEmployee) GetGroups() []HrisGroup {
	if o == nil {
		return nil
	}
	return o.Groups
}

func (o *HrisEmployee) GetHiredAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.HiredAt
}

func (o *HrisEmployee) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *HrisEmployee) GetImageURL() *string {
	if o == nil {
		return nil
	}
	return o.ImageURL
}

func (o *HrisEmployee) GetLanguageLocale() *string {
	if o == nil {
		return nil
	}
	return o.LanguageLocale
}

func (o *HrisEmployee) GetLocation() *string {
	if o == nil {
		return nil
	}
	return o.Location
}

func (o *HrisEmployee) GetLocations() []HrisLocation {
	if o == nil {
		return nil
	}
	return o.Locations
}

func (o *HrisEmployee) GetManagerID() *string {
	if o == nil {
		return nil
	}
	return o.ManagerID
}

func (o *HrisEmployee) GetMaritalStatus() *MaritalStatus {
	if o == nil {
		return nil
	}
	return o.MaritalStatus
}

func (o *HrisEmployee) GetMetadata() []HrisMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *HrisEmployee) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *HrisEmployee) GetPronouns() *string {
	if o == nil {
		return nil
	}
	return o.Pronouns
}

func (o *HrisEmployee) GetRaw() *HrisEmployeeRaw {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *HrisEmployee) GetSalutation() *string {
	if o == nil {
		return nil
	}
	return o.Salutation
}

func (o *HrisEmployee) GetSsnSin() *string {
	if o == nil {
		return nil
	}
	return o.SsnSin
}

func (o *HrisEmployee) GetTelephones() []HrisTelephone {
	if o == nil {
		return nil
	}
	return o.Telephones
}

func (o *HrisEmployee) GetTerminatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.TerminatedAt
}

func (o *HrisEmployee) GetTimezone() *string {
	if o == nil {
		return nil
	}
	return o.Timezone
}

func (o *HrisEmployee) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *HrisEmployee) GetUpdatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

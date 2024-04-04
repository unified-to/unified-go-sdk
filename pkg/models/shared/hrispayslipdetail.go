// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type HrisPayslipDetailType string

const (
	HrisPayslipDetailTypeEarningSalary                  HrisPayslipDetailType = "EARNING_SALARY"
	HrisPayslipDetailTypeEarningOvertime                HrisPayslipDetailType = "EARNING_OVERTIME"
	HrisPayslipDetailTypeEarningTip                     HrisPayslipDetailType = "EARNING_TIP"
	HrisPayslipDetailTypeEarningBonus                   HrisPayslipDetailType = "EARNING_BONUS"
	HrisPayslipDetailTypeEarningCommission              HrisPayslipDetailType = "EARNING_COMMISSION"
	HrisPayslipDetailTypeEarningAdjustment              HrisPayslipDetailType = "EARNING_ADJUSTMENT"
	HrisPayslipDetailTypeEarning                        HrisPayslipDetailType = "EARNING"
	HrisPayslipDetailTypePretaxDeduction                HrisPayslipDetailType = "PRETAX_DEDUCTION"
	HrisPayslipDetailTypePretaxDeductionHealthInsurance HrisPayslipDetailType = "PRETAX_DEDUCTION_HEALTH_INSURANCE"
	HrisPayslipDetailTypePretaxDeductionRetirement      HrisPayslipDetailType = "PRETAX_DEDUCTION_RETIREMENT"
	HrisPayslipDetailTypePretaxDeductionHra             HrisPayslipDetailType = "PRETAX_DEDUCTION_HRA"
	HrisPayslipDetailTypeTaxFederal                     HrisPayslipDetailType = "TAX_FEDERAL"
	HrisPayslipDetailTypeTaxRegion                      HrisPayslipDetailType = "TAX_REGION"
	HrisPayslipDetailTypeTaxLocal                       HrisPayslipDetailType = "TAX_LOCAL"
	HrisPayslipDetailTypePosttaxBenefit                 HrisPayslipDetailType = "POSTTAX_BENEFIT"
	HrisPayslipDetailTypePosttaxGarnishment             HrisPayslipDetailType = "POSTTAX_GARNISHMENT"
	HrisPayslipDetailTypeReimbursement                  HrisPayslipDetailType = "REIMBURSEMENT"
)

func (e HrisPayslipDetailType) ToPointer() *HrisPayslipDetailType {
	return &e
}

func (e *HrisPayslipDetailType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "EARNING_SALARY":
		fallthrough
	case "EARNING_OVERTIME":
		fallthrough
	case "EARNING_TIP":
		fallthrough
	case "EARNING_BONUS":
		fallthrough
	case "EARNING_COMMISSION":
		fallthrough
	case "EARNING_ADJUSTMENT":
		fallthrough
	case "EARNING":
		fallthrough
	case "PRETAX_DEDUCTION":
		fallthrough
	case "PRETAX_DEDUCTION_HEALTH_INSURANCE":
		fallthrough
	case "PRETAX_DEDUCTION_RETIREMENT":
		fallthrough
	case "PRETAX_DEDUCTION_HRA":
		fallthrough
	case "TAX_FEDERAL":
		fallthrough
	case "TAX_REGION":
		fallthrough
	case "TAX_LOCAL":
		fallthrough
	case "POSTTAX_BENEFIT":
		fallthrough
	case "POSTTAX_GARNISHMENT":
		fallthrough
	case "REIMBURSEMENT":
		*e = HrisPayslipDetailType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for HrisPayslipDetailType: %v", v)
	}
}

type HrisPayslipDetail struct {
	Amount         float64                `json:"amount"`
	CompanyAmount  *float64               `json:"company_amount,omitempty"`
	Description    *string                `json:"description,omitempty"`
	EmployeeAmount *float64               `json:"employee_amount,omitempty"`
	Name           *string                `json:"name,omitempty"`
	Type           *HrisPayslipDetailType `json:"type,omitempty"`
}

func (o *HrisPayslipDetail) GetAmount() float64 {
	if o == nil {
		return 0.0
	}
	return o.Amount
}

func (o *HrisPayslipDetail) GetCompanyAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.CompanyAmount
}

func (o *HrisPayslipDetail) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *HrisPayslipDetail) GetEmployeeAmount() *float64 {
	if o == nil {
		return nil
	}
	return o.EmployeeAmount
}

func (o *HrisPayslipDetail) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *HrisPayslipDetail) GetType() *HrisPayslipDetailType {
	if o == nil {
		return nil
	}
	return o.Type
}

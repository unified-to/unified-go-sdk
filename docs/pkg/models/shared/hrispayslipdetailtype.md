# HrisPayslipDetailType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HrisPayslipDetailTypeEarningSalary

// Open enum: custom values can be created with a direct type cast
custom := shared.HrisPayslipDetailType("custom_value")
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `HrisPayslipDetailTypeEarningSalary`                  | EARNING_SALARY                                        |
| `HrisPayslipDetailTypeEarningOvertime`                | EARNING_OVERTIME                                      |
| `HrisPayslipDetailTypeEarningTip`                     | EARNING_TIP                                           |
| `HrisPayslipDetailTypeEarningBonus`                   | EARNING_BONUS                                         |
| `HrisPayslipDetailTypeEarningCommission`              | EARNING_COMMISSION                                    |
| `HrisPayslipDetailTypeEarningAdjustment`              | EARNING_ADJUSTMENT                                    |
| `HrisPayslipDetailTypeEarning`                        | EARNING                                               |
| `HrisPayslipDetailTypePretaxDeduction`                | PRETAX_DEDUCTION                                      |
| `HrisPayslipDetailTypePretaxDeductionHealthInsurance` | PRETAX_DEDUCTION_HEALTH_INSURANCE                     |
| `HrisPayslipDetailTypePretaxDeductionRetirement`      | PRETAX_DEDUCTION_RETIREMENT                           |
| `HrisPayslipDetailTypePretaxDeductionHra`             | PRETAX_DEDUCTION_HRA                                  |
| `HrisPayslipDetailTypeTaxFederal`                     | TAX_FEDERAL                                           |
| `HrisPayslipDetailTypeTaxRegion`                      | TAX_REGION                                            |
| `HrisPayslipDetailTypeTaxLocal`                       | TAX_LOCAL                                             |
| `HrisPayslipDetailTypePosttaxBenefit`                 | POSTTAX_BENEFIT                                       |
| `HrisPayslipDetailTypePosttaxGarnishment`             | POSTTAX_GARNISHMENT                                   |
| `HrisPayslipDetailTypeReimbursement`                  | REIMBURSEMENT                                         |
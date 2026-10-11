# DetailType

the normalized payslip-line type this code produces

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.DetailTypeEarningSalary

// Open enum: custom values can be created with a direct type cast
custom := shared.DetailType("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `DetailTypeEarningSalary`                  | EARNING_SALARY                             |
| `DetailTypeEarningOvertime`                | EARNING_OVERTIME                           |
| `DetailTypeEarningTip`                     | EARNING_TIP                                |
| `DetailTypeEarningBonus`                   | EARNING_BONUS                              |
| `DetailTypeEarningCommission`              | EARNING_COMMISSION                         |
| `DetailTypeEarningAdjustment`              | EARNING_ADJUSTMENT                         |
| `DetailTypeEarning`                        | EARNING                                    |
| `DetailTypePretaxDeduction`                | PRETAX_DEDUCTION                           |
| `DetailTypePretaxDeductionHealthInsurance` | PRETAX_DEDUCTION_HEALTH_INSURANCE          |
| `DetailTypePretaxDeductionRetirement`      | PRETAX_DEDUCTION_RETIREMENT                |
| `DetailTypePretaxDeductionHra`             | PRETAX_DEDUCTION_HRA                       |
| `DetailTypeTaxFederal`                     | TAX_FEDERAL                                |
| `DetailTypeTaxRegion`                      | TAX_REGION                                 |
| `DetailTypeTaxLocal`                       | TAX_LOCAL                                  |
| `DetailTypePosttaxBenefit`                 | POSTTAX_BENEFIT                            |
| `DetailTypePosttaxGarnishment`             | POSTTAX_GARNISHMENT                        |
| `DetailTypeReimbursement`                  | REIMBURSEMENT                              |
| `DetailTypeEarningRegular`                 | EARNING_REGULAR                            |
| `DetailTypeEarningDoubleOvertime`          | EARNING_DOUBLE_OVERTIME                    |
| `DetailTypeEarningPto`                     | EARNING_PTO                                |
| `DetailTypeEarningSick`                    | EARNING_SICK                               |
| `DetailTypeEarningHoliday`                 | EARNING_HOLIDAY                            |
| `DetailTypeEarningSeverance`               | EARNING_SEVERANCE                          |
| `DetailTypeEarningReimbursement`           | EARNING_REIMBURSEMENT                      |
| `DetailTypeTaxSocialSecurity`              | TAX_SOCIAL_SECURITY                        |
| `DetailTypeTaxMedicare`                    | TAX_MEDICARE                               |
| `DetailTypeTaxUnemployment`                | TAX_UNEMPLOYMENT                           |
| `DetailTypeTaxEmployer`                    | TAX_EMPLOYER                               |
| `DetailTypePretaxDeductionHsa`             | PRETAX_DEDUCTION_HSA                       |
| `DetailTypePretaxDeductionFsa`             | PRETAX_DEDUCTION_FSA                       |
| `DetailTypePretaxDeductionDependentCare`   | PRETAX_DEDUCTION_DEPENDENT_CARE            |
| `DetailTypePretaxDeductionCommuter`        | PRETAX_DEDUCTION_COMMUTER                  |
| `DetailTypePretaxDeductionDental`          | PRETAX_DEDUCTION_DENTAL                    |
| `DetailTypePretaxDeductionVision`          | PRETAX_DEDUCTION_VISION                    |
| `DetailTypePosttaxDeduction`               | POSTTAX_DEDUCTION                          |
| `DetailTypePosttaxRetirementRoth`          | POSTTAX_RETIREMENT_ROTH                    |
| `DetailTypePosttaxLoanRepayment`           | POSTTAX_LOAN_REPAYMENT                     |
| `DetailTypeEmployerContribution`           | EMPLOYER_CONTRIBUTION                      |
| `DetailTypePension`                        | PENSION                                    |
| `DetailTypeEarningCashInLieu`              | EARNING_CASH_IN_LIEU                       |
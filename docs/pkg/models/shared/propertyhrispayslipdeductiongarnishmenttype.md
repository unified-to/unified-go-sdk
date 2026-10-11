# PropertyHrisPayslipDeductionGarnishmentType

set when this deduction is a court-ordered garnishment

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PropertyHrisPayslipDeductionGarnishmentTypeChildSupport

// Open enum: custom values can be created with a direct type cast
custom := shared.PropertyHrisPayslipDeductionGarnishmentType("custom_value")
```


## Values

| Name                                                      | Value                                                     |
| --------------------------------------------------------- | --------------------------------------------------------- |
| `PropertyHrisPayslipDeductionGarnishmentTypeChildSupport` | CHILD_SUPPORT                                             |
| `PropertyHrisPayslipDeductionGarnishmentTypeTaxLevy`      | TAX_LEVY                                                  |
| `PropertyHrisPayslipDeductionGarnishmentTypeStudentLoan`  | STUDENT_LOAN                                              |
| `PropertyHrisPayslipDeductionGarnishmentTypeCreditor`     | CREDITOR                                                  |
| `PropertyHrisPayslipDeductionGarnishmentTypeBankruptcy`   | BANKRUPTCY                                                |
| `PropertyHrisPayslipDeductionGarnishmentTypeOther`        | OTHER                                                     |
# GarnishmentType

set when this deduction is a court-ordered garnishment

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.GarnishmentTypeChildSupport

// Open enum: custom values can be created with a direct type cast
custom := shared.GarnishmentType("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `GarnishmentTypeChildSupport` | CHILD_SUPPORT                 |
| `GarnishmentTypeTaxLevy`      | TAX_LEVY                      |
| `GarnishmentTypeStudentLoan`  | STUDENT_LOAN                  |
| `GarnishmentTypeCreditor`     | CREDITOR                      |
| `GarnishmentTypeBankruptcy`   | BANKRUPTCY                    |
| `GarnishmentTypeOther`        | OTHER                         |
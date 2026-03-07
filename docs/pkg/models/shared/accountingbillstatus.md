# AccountingBillStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AccountingBillStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := shared.AccountingBillStatus("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `AccountingBillStatusDraft`             | DRAFT                                   |
| `AccountingBillStatusVoided`            | VOIDED                                  |
| `AccountingBillStatusAuthorized`        | AUTHORIZED                              |
| `AccountingBillStatusPaid`              | PAID                                    |
| `AccountingBillStatusPartiallyPaid`     | PARTIALLY_PAID                          |
| `AccountingBillStatusPartiallyRefunded` | PARTIALLY_REFUNDED                      |
| `AccountingBillStatusRefunded`          | REFUNDED                                |
| `AccountingBillStatusSubmitted`         | SUBMITTED                               |
| `AccountingBillStatusDeleted`           | DELETED                                 |
| `AccountingBillStatusOverdue`           | OVERDUE                                 |
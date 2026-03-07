# AccountingCreditmemoStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AccountingCreditmemoStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := shared.AccountingCreditmemoStatus("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `AccountingCreditmemoStatusDraft`             | DRAFT                                         |
| `AccountingCreditmemoStatusVoided`            | VOIDED                                        |
| `AccountingCreditmemoStatusAuthorized`        | AUTHORIZED                                    |
| `AccountingCreditmemoStatusPaid`              | PAID                                          |
| `AccountingCreditmemoStatusPartiallyPaid`     | PARTIALLY_PAID                                |
| `AccountingCreditmemoStatusPartiallyRefunded` | PARTIALLY_REFUNDED                            |
| `AccountingCreditmemoStatusRefunded`          | REFUNDED                                      |
| `AccountingCreditmemoStatusSubmitted`         | SUBMITTED                                     |
| `AccountingCreditmemoStatusDeleted`           | DELETED                                       |
| `AccountingCreditmemoStatusOverdue`           | OVERDUE                                       |
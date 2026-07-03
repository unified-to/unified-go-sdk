# AccountingVendorcreditStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AccountingVendorcreditStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := shared.AccountingVendorcreditStatus("custom_value")
```


## Values

| Name                                            | Value                                           |
| ----------------------------------------------- | ----------------------------------------------- |
| `AccountingVendorcreditStatusDraft`             | DRAFT                                           |
| `AccountingVendorcreditStatusVoided`            | VOIDED                                          |
| `AccountingVendorcreditStatusAuthorized`        | AUTHORIZED                                      |
| `AccountingVendorcreditStatusPaid`              | PAID                                            |
| `AccountingVendorcreditStatusPartiallyPaid`     | PARTIALLY_PAID                                  |
| `AccountingVendorcreditStatusPartiallyRefunded` | PARTIALLY_REFUNDED                              |
| `AccountingVendorcreditStatusRefunded`          | REFUNDED                                        |
| `AccountingVendorcreditStatusSubmitted`         | SUBMITTED                                       |
| `AccountingVendorcreditStatusDeleted`           | DELETED                                         |
| `AccountingVendorcreditStatusOverdue`           | OVERDUE                                         |
# AccountingInvoiceStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AccountingInvoiceStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := shared.AccountingInvoiceStatus("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `AccountingInvoiceStatusDraft`             | DRAFT                                      |
| `AccountingInvoiceStatusVoided`            | VOIDED                                     |
| `AccountingInvoiceStatusAuthorized`        | AUTHORIZED                                 |
| `AccountingInvoiceStatusPaid`              | PAID                                       |
| `AccountingInvoiceStatusPartiallyPaid`     | PARTIALLY_PAID                             |
| `AccountingInvoiceStatusPartiallyRefunded` | PARTIALLY_REFUNDED                         |
| `AccountingInvoiceStatusRefunded`          | REFUNDED                                   |
| `AccountingInvoiceStatusSubmitted`         | SUBMITTED                                  |
| `AccountingInvoiceStatusDeleted`           | DELETED                                    |
| `AccountingInvoiceStatusOverdue`           | OVERDUE                                    |
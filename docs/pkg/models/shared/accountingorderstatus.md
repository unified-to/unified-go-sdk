# AccountingOrderStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AccountingOrderStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := shared.AccountingOrderStatus("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `AccountingOrderStatusDraft`             | DRAFT                                    |
| `AccountingOrderStatusVoided`            | VOIDED                                   |
| `AccountingOrderStatusAuthorized`        | AUTHORIZED                               |
| `AccountingOrderStatusPaid`              | PAID                                     |
| `AccountingOrderStatusPartiallyPaid`     | PARTIALLY_PAID                           |
| `AccountingOrderStatusPartiallyRefunded` | PARTIALLY_REFUNDED                       |
| `AccountingOrderStatusRefunded`          | REFUNDED                                 |
| `AccountingOrderStatusSubmitted`         | SUBMITTED                                |
| `AccountingOrderStatusDeleted`           | DELETED                                  |
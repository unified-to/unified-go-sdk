# AccountingPurchaseorderStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AccountingPurchaseorderStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := shared.AccountingPurchaseorderStatus("custom_value")
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `AccountingPurchaseorderStatusDraft`             | DRAFT                                            |
| `AccountingPurchaseorderStatusVoided`            | VOIDED                                           |
| `AccountingPurchaseorderStatusAuthorized`        | AUTHORIZED                                       |
| `AccountingPurchaseorderStatusPaid`              | PAID                                             |
| `AccountingPurchaseorderStatusPartiallyPaid`     | PARTIALLY_PAID                                   |
| `AccountingPurchaseorderStatusPartiallyRefunded` | PARTIALLY_REFUNDED                               |
| `AccountingPurchaseorderStatusRefunded`          | REFUNDED                                         |
| `AccountingPurchaseorderStatusSubmitted`         | SUBMITTED                                        |
| `AccountingPurchaseorderStatusDeleted`           | DELETED                                          |
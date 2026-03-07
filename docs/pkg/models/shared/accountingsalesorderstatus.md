# AccountingSalesorderStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AccountingSalesorderStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := shared.AccountingSalesorderStatus("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `AccountingSalesorderStatusDraft`             | DRAFT                                         |
| `AccountingSalesorderStatusVoided`            | VOIDED                                        |
| `AccountingSalesorderStatusAuthorized`        | AUTHORIZED                                    |
| `AccountingSalesorderStatusPaid`              | PAID                                          |
| `AccountingSalesorderStatusPartiallyPaid`     | PARTIALLY_PAID                                |
| `AccountingSalesorderStatusPartiallyRefunded` | PARTIALLY_REFUNDED                            |
| `AccountingSalesorderStatusRefunded`          | REFUNDED                                      |
| `AccountingSalesorderStatusSubmitted`         | SUBMITTED                                     |
| `AccountingSalesorderStatusDeleted`           | DELETED                                       |
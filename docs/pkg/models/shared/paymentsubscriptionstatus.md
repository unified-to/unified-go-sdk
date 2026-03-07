# PaymentSubscriptionStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PaymentSubscriptionStatusActive

// Open enum: custom values can be created with a direct type cast
custom := shared.PaymentSubscriptionStatus("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `PaymentSubscriptionStatusActive`   | ACTIVE                              |
| `PaymentSubscriptionStatusInactive` | INACTIVE                            |
| `PaymentSubscriptionStatusCanceled` | CANCELED                            |
| `PaymentSubscriptionStatusPaused`   | PAUSED                              |
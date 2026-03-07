# PaymentPayoutStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PaymentPayoutStatusSucceeded

// Open enum: custom values can be created with a direct type cast
custom := shared.PaymentPayoutStatus("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `PaymentPayoutStatusSucceeded` | SUCCEEDED                      |
| `PaymentPayoutStatusPending`   | PENDING                        |
| `PaymentPayoutStatusFailed`    | FAILED                         |
| `PaymentPayoutStatusCanceled`  | CANCELED                       |
# PaymentPaymentStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PaymentPaymentStatusSucceeded

// Open enum: custom values can be created with a direct type cast
custom := shared.PaymentPaymentStatus("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `PaymentPaymentStatusSucceeded`  | SUCCEEDED                        |
| `PaymentPaymentStatusPending`    | PENDING                          |
| `PaymentPaymentStatusAuthorized` | AUTHORIZED                       |
| `PaymentPaymentStatusFailed`     | FAILED                           |
| `PaymentPaymentStatusCanceled`   | CANCELED                         |
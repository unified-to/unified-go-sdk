# PaymentRefundStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PaymentRefundStatusSucceeded

// Open enum: custom values can be created with a direct type cast
custom := shared.PaymentRefundStatus("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `PaymentRefundStatusSucceeded` | SUCCEEDED                      |
| `PaymentRefundStatusPending`   | PENDING                        |
| `PaymentRefundStatusFailed`    | FAILED                         |
| `PaymentRefundStatusCanceled`  | CANCELED                       |
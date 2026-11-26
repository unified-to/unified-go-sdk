# PrintingOrderStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PrintingOrderStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := shared.PrintingOrderStatus("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `PrintingOrderStatusDraft`            | DRAFT                                 |
| `PrintingOrderStatusPending`          | PENDING                               |
| `PrintingOrderStatusOnHold`           | ON_HOLD                               |
| `PrintingOrderStatusInProduction`     | IN_PRODUCTION                         |
| `PrintingOrderStatusPartiallyShipped` | PARTIALLY_SHIPPED                     |
| `PrintingOrderStatusShipped`          | SHIPPED                               |
| `PrintingOrderStatusDelivered`        | DELIVERED                             |
| `PrintingOrderStatusCancelled`        | CANCELLED                             |
| `PrintingOrderStatusFailed`           | FAILED                                |
| `PrintingOrderStatusOther`            | OTHER                                 |
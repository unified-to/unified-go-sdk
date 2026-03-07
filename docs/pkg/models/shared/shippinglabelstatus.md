# ShippingLabelStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ShippingLabelStatusPending

// Open enum: custom values can be created with a direct type cast
custom := shared.ShippingLabelStatus("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `ShippingLabelStatusPending`           | PENDING                                |
| `ShippingLabelStatusProcessing`        | PROCESSING                             |
| `ShippingLabelStatusInTransit`         | IN_TRANSIT                             |
| `ShippingLabelStatusDelivered`         | DELIVERED                              |
| `ShippingLabelStatusException`         | EXCEPTION                              |
| `ShippingLabelStatusCancelled`         | CANCELLED                              |
| `ShippingLabelStatusLabelCreated`      | LABEL_CREATED                          |
| `ShippingLabelStatusPickedUp`          | PICKED_UP                              |
| `ShippingLabelStatusOutForDelivery`    | OUT_FOR_DELIVERY                       |
| `ShippingLabelStatusDeliveryAttempted` | DELIVERY_ATTEMPTED                     |
| `ShippingLabelStatusReturnedToSender`  | RETURNED_TO_SENDER                     |
| `ShippingLabelStatusHeldAtLocation`    | HELD_AT_LOCATION                       |
| `ShippingLabelStatusCustomsClearance`  | CUSTOMS_CLEARANCE                      |
| `ShippingLabelStatusExceptionResolved` | EXCEPTION_RESOLVED                     |
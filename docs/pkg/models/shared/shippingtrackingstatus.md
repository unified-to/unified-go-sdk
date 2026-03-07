# ShippingTrackingStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ShippingTrackingStatusPending

// Open enum: custom values can be created with a direct type cast
custom := shared.ShippingTrackingStatus("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `ShippingTrackingStatusPending`           | PENDING                                   |
| `ShippingTrackingStatusProcessing`        | PROCESSING                                |
| `ShippingTrackingStatusInTransit`         | IN_TRANSIT                                |
| `ShippingTrackingStatusDelivered`         | DELIVERED                                 |
| `ShippingTrackingStatusException`         | EXCEPTION                                 |
| `ShippingTrackingStatusCancelled`         | CANCELLED                                 |
| `ShippingTrackingStatusLabelCreated`      | LABEL_CREATED                             |
| `ShippingTrackingStatusPickedUp`          | PICKED_UP                                 |
| `ShippingTrackingStatusOutForDelivery`    | OUT_FOR_DELIVERY                          |
| `ShippingTrackingStatusDeliveryAttempted` | DELIVERY_ATTEMPTED                        |
| `ShippingTrackingStatusReturnedToSender`  | RETURNED_TO_SENDER                        |
| `ShippingTrackingStatusHeldAtLocation`    | HELD_AT_LOCATION                          |
| `ShippingTrackingStatusCustomsClearance`  | CUSTOMS_CLEARANCE                         |
| `ShippingTrackingStatusExceptionResolved` | EXCEPTION_RESOLVED                        |
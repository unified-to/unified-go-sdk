# ShippingTrackingEventStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ShippingTrackingEventStatusPending

// Open enum: custom values can be created with a direct type cast
custom := shared.ShippingTrackingEventStatus("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `ShippingTrackingEventStatusPending`           | PENDING                                        |
| `ShippingTrackingEventStatusProcessing`        | PROCESSING                                     |
| `ShippingTrackingEventStatusInTransit`         | IN_TRANSIT                                     |
| `ShippingTrackingEventStatusDelivered`         | DELIVERED                                      |
| `ShippingTrackingEventStatusException`         | EXCEPTION                                      |
| `ShippingTrackingEventStatusCancelled`         | CANCELLED                                      |
| `ShippingTrackingEventStatusLabelCreated`      | LABEL_CREATED                                  |
| `ShippingTrackingEventStatusPickedUp`          | PICKED_UP                                      |
| `ShippingTrackingEventStatusOutForDelivery`    | OUT_FOR_DELIVERY                               |
| `ShippingTrackingEventStatusDeliveryAttempted` | DELIVERY_ATTEMPTED                             |
| `ShippingTrackingEventStatusReturnedToSender`  | RETURNED_TO_SENDER                             |
| `ShippingTrackingEventStatusHeldAtLocation`    | HELD_AT_LOCATION                               |
| `ShippingTrackingEventStatusCustomsClearance`  | CUSTOMS_CLEARANCE                              |
| `ShippingTrackingEventStatusExceptionResolved` | EXCEPTION_RESOLVED                             |
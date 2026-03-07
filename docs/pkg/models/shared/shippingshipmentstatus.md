# ShippingShipmentStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ShippingShipmentStatusPending

// Open enum: custom values can be created with a direct type cast
custom := shared.ShippingShipmentStatus("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `ShippingShipmentStatusPending`           | PENDING                                   |
| `ShippingShipmentStatusProcessing`        | PROCESSING                                |
| `ShippingShipmentStatusInTransit`         | IN_TRANSIT                                |
| `ShippingShipmentStatusDelivered`         | DELIVERED                                 |
| `ShippingShipmentStatusException`         | EXCEPTION                                 |
| `ShippingShipmentStatusCancelled`         | CANCELLED                                 |
| `ShippingShipmentStatusLabelCreated`      | LABEL_CREATED                             |
| `ShippingShipmentStatusPickedUp`          | PICKED_UP                                 |
| `ShippingShipmentStatusOutForDelivery`    | OUT_FOR_DELIVERY                          |
| `ShippingShipmentStatusDeliveryAttempted` | DELIVERY_ATTEMPTED                        |
| `ShippingShipmentStatusReturnedToSender`  | RETURNED_TO_SENDER                        |
| `ShippingShipmentStatusHeldAtLocation`    | HELD_AT_LOCATION                          |
| `ShippingShipmentStatusCustomsClearance`  | CUSTOMS_CLEARANCE                         |
| `ShippingShipmentStatusExceptionResolved` | EXCEPTION_RESOLVED                        |
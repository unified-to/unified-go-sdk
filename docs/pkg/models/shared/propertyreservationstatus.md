# PropertyReservationStatus

the reservation status

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PropertyReservationStatusInquiry

// Open enum: custom values can be created with a direct type cast
custom := shared.PropertyReservationStatus("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `PropertyReservationStatusInquiry`    | INQUIRY                               |
| `PropertyReservationStatusPending`    | PENDING                               |
| `PropertyReservationStatusConfirmed`  | CONFIRMED                             |
| `PropertyReservationStatusCancelled`  | CANCELLED                             |
| `PropertyReservationStatusCheckedIn`  | CHECKED_IN                            |
| `PropertyReservationStatusCheckedOut` | CHECKED_OUT                           |
| `PropertyReservationStatusNoShow`     | NO_SHOW                               |
| `PropertyReservationStatusOther`      | OTHER                                 |
# DmsAppointmentStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.DmsAppointmentStatusScheduled

// Open enum: custom values can be created with a direct type cast
custom := shared.DmsAppointmentStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `DmsAppointmentStatusScheduled` | SCHEDULED                       |
| `DmsAppointmentStatusConfirmed` | CONFIRMED                       |
| `DmsAppointmentStatusConverted` | CONVERTED                       |
| `DmsAppointmentStatusCancelled` | CANCELLED                       |
| `DmsAppointmentStatusNoShow`    | NO_SHOW                         |
| `DmsAppointmentStatusOther`     | OTHER                           |
# DmsVehiclesaleStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.DmsVehiclesaleStatusQuote

// Open enum: custom values can be created with a direct type cast
custom := shared.DmsVehiclesaleStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `DmsVehiclesaleStatusQuote`     | QUOTE                           |
| `DmsVehiclesaleStatusPending`   | PENDING                         |
| `DmsVehiclesaleStatusSold`      | SOLD                            |
| `DmsVehiclesaleStatusDelivered` | DELIVERED                       |
| `DmsVehiclesaleStatusFinalized` | FINALIZED                       |
| `DmsVehiclesaleStatusCancelled` | CANCELLED                       |
| `DmsVehiclesaleStatusOther`     | OTHER                           |
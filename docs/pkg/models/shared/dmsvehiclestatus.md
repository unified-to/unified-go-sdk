# DmsVehicleStatus

Stock status; only set for inventory units

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.DmsVehicleStatusAvailable

// Open enum: custom values can be created with a direct type cast
custom := shared.DmsVehicleStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `DmsVehicleStatusAvailable`   | AVAILABLE                     |
| `DmsVehicleStatusOnOrder`     | ON_ORDER                      |
| `DmsVehicleStatusInTransit`   | IN_TRANSIT                    |
| `DmsVehicleStatusReserved`    | RESERVED                      |
| `DmsVehicleStatusSold`        | SOLD                          |
| `DmsVehicleStatusUnavailable` | UNAVAILABLE                   |
| `DmsVehicleStatusOther`       | OTHER                         |
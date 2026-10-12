# HealthcareSlotStatus

Availability of the slot

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HealthcareSlotStatusFree

// Open enum: custom values can be created with a direct type cast
custom := shared.HealthcareSlotStatus("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `HealthcareSlotStatusFree`            | FREE                                  |
| `HealthcareSlotStatusBusy`            | BUSY                                  |
| `HealthcareSlotStatusBusyUnavailable` | BUSY_UNAVAILABLE                      |
| `HealthcareSlotStatusBusyTentative`   | BUSY_TENTATIVE                        |
| `HealthcareSlotStatusEnteredInError`  | ENTERED_IN_ERROR                      |
| `HealthcareSlotStatusOther`           | OTHER                                 |
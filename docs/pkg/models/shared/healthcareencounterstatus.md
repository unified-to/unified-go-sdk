# HealthcareEncounterStatus

Status of the encounter

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HealthcareEncounterStatusPlanned

// Open enum: custom values can be created with a direct type cast
custom := shared.HealthcareEncounterStatus("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `HealthcareEncounterStatusPlanned`        | PLANNED                                   |
| `HealthcareEncounterStatusArrived`        | ARRIVED                                   |
| `HealthcareEncounterStatusTriaged`        | TRIAGED                                   |
| `HealthcareEncounterStatusInProgress`     | IN_PROGRESS                               |
| `HealthcareEncounterStatusOnleave`        | ONLEAVE                                   |
| `HealthcareEncounterStatusFinished`       | FINISHED                                  |
| `HealthcareEncounterStatusCancelled`      | CANCELLED                                 |
| `HealthcareEncounterStatusEnteredInError` | ENTERED_IN_ERROR                          |
| `HealthcareEncounterStatusUnknown`        | UNKNOWN                                   |
| `HealthcareEncounterStatusOther`          | OTHER                                     |
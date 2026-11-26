# FhirEncounterStatus

planned | arrived | triaged | in-progress | onleave | finished | cancelled | entered-in-error | unknown

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.FhirEncounterStatusPlanned

// Open enum: custom values can be created with a direct type cast
custom := shared.FhirEncounterStatus("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `FhirEncounterStatusPlanned`        | planned                             |
| `FhirEncounterStatusArrived`        | arrived                             |
| `FhirEncounterStatusTriaged`        | triaged                             |
| `FhirEncounterStatusInProgress`     | in-progress                         |
| `FhirEncounterStatusOnleave`        | onleave                             |
| `FhirEncounterStatusFinished`       | finished                            |
| `FhirEncounterStatusCancelled`      | cancelled                           |
| `FhirEncounterStatusEnteredInError` | entered-in-error                    |
| `FhirEncounterStatusUnknown`        | unknown                             |
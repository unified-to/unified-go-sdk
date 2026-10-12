# HealthcareObservationStatus

Status of the observation result

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HealthcareObservationStatusRegistered

// Open enum: custom values can be created with a direct type cast
custom := shared.HealthcareObservationStatus("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `HealthcareObservationStatusRegistered`     | REGISTERED                                  |
| `HealthcareObservationStatusPreliminary`    | PRELIMINARY                                 |
| `HealthcareObservationStatusFinal`          | FINAL                                       |
| `HealthcareObservationStatusAmended`        | AMENDED                                     |
| `HealthcareObservationStatusCorrected`      | CORRECTED                                   |
| `HealthcareObservationStatusCancelled`      | CANCELLED                                   |
| `HealthcareObservationStatusEnteredInError` | ENTERED_IN_ERROR                            |
| `HealthcareObservationStatusUnknown`        | UNKNOWN                                     |
| `HealthcareObservationStatusOther`          | OTHER                                       |
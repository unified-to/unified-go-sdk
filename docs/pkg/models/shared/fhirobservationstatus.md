# FhirObservationStatus

registered | preliminary | final | amended | corrected | cancelled | entered-in-error | unknown

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.FhirObservationStatusRegistered

// Open enum: custom values can be created with a direct type cast
custom := shared.FhirObservationStatus("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `FhirObservationStatusRegistered`     | registered                            |
| `FhirObservationStatusPreliminary`    | preliminary                           |
| `FhirObservationStatusFinal`          | final                                 |
| `FhirObservationStatusAmended`        | amended                               |
| `FhirObservationStatusCorrected`      | corrected                             |
| `FhirObservationStatusCancelled`      | cancelled                             |
| `FhirObservationStatusEnteredInError` | entered-in-error                      |
| `FhirObservationStatusUnknown`        | unknown                               |
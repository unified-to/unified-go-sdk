# FhirMedicationRequestStatus

active | on-hold | cancelled | completed | entered-in-error | stopped | draft | unknown

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.FhirMedicationRequestStatusActive

// Open enum: custom values can be created with a direct type cast
custom := shared.FhirMedicationRequestStatus("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `FhirMedicationRequestStatusActive`         | active                                      |
| `FhirMedicationRequestStatusOnHold`         | on-hold                                     |
| `FhirMedicationRequestStatusCancelled`      | cancelled                                   |
| `FhirMedicationRequestStatusCompleted`      | completed                                   |
| `FhirMedicationRequestStatusEnteredInError` | entered-in-error                            |
| `FhirMedicationRequestStatusStopped`        | stopped                                     |
| `FhirMedicationRequestStatusDraft`          | draft                                       |
| `FhirMedicationRequestStatusUnknown`        | unknown                                     |
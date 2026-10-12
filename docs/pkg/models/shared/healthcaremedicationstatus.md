# HealthcareMedicationStatus

Status of the medication order

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HealthcareMedicationStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := shared.HealthcareMedicationStatus("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `HealthcareMedicationStatusDraft`          | DRAFT                                      |
| `HealthcareMedicationStatusActive`         | ACTIVE                                     |
| `HealthcareMedicationStatusOnHold`         | ON_HOLD                                    |
| `HealthcareMedicationStatusCompleted`      | COMPLETED                                  |
| `HealthcareMedicationStatusStopped`        | STOPPED                                    |
| `HealthcareMedicationStatusCancelled`      | CANCELLED                                  |
| `HealthcareMedicationStatusEnteredInError` | ENTERED_IN_ERROR                           |
| `HealthcareMedicationStatusUnknown`        | UNKNOWN                                    |
| `HealthcareMedicationStatusOther`          | OTHER                                      |
# HealthcareCareplanStatus

Status of the care plan

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HealthcareCareplanStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := shared.HealthcareCareplanStatus("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `HealthcareCareplanStatusDraft`          | DRAFT                                    |
| `HealthcareCareplanStatusActive`         | ACTIVE                                   |
| `HealthcareCareplanStatusOnHold`         | ON_HOLD                                  |
| `HealthcareCareplanStatusRevoked`        | REVOKED                                  |
| `HealthcareCareplanStatusCompleted`      | COMPLETED                                |
| `HealthcareCareplanStatusEnteredInError` | ENTERED_IN_ERROR                         |
| `HealthcareCareplanStatusUnknown`        | UNKNOWN                                  |
| `HealthcareCareplanStatusOther`          | OTHER                                    |
# HealthcareGoalStatus

Lifecycle status of the goal

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HealthcareGoalStatusProposed

// Open enum: custom values can be created with a direct type cast
custom := shared.HealthcareGoalStatus("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `HealthcareGoalStatusProposed`       | PROPOSED                             |
| `HealthcareGoalStatusPlanned`        | PLANNED                              |
| `HealthcareGoalStatusAccepted`       | ACCEPTED                             |
| `HealthcareGoalStatusActive`         | ACTIVE                               |
| `HealthcareGoalStatusOnHold`         | ON_HOLD                              |
| `HealthcareGoalStatusCompleted`      | COMPLETED                            |
| `HealthcareGoalStatusCancelled`      | CANCELLED                            |
| `HealthcareGoalStatusEnteredInError` | ENTERED_IN_ERROR                     |
| `HealthcareGoalStatusRejected`       | REJECTED                             |
| `HealthcareGoalStatusOther`          | OTHER                                |
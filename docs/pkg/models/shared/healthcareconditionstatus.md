# HealthcareConditionStatus

Clinical status of the condition

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HealthcareConditionStatusActive

// Open enum: custom values can be created with a direct type cast
custom := shared.HealthcareConditionStatus("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `HealthcareConditionStatusActive`     | ACTIVE                                |
| `HealthcareConditionStatusRecurrence` | RECURRENCE                            |
| `HealthcareConditionStatusRelapse`    | RELAPSE                               |
| `HealthcareConditionStatusInactive`   | INACTIVE                              |
| `HealthcareConditionStatusRemission`  | REMISSION                             |
| `HealthcareConditionStatusResolved`   | RESOLVED                              |
| `HealthcareConditionStatusOther`      | OTHER                                 |
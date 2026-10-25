# HealthcareProcedureStatus

Status of the procedure

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.HealthcareProcedureStatusPreparation

// Open enum: custom values can be created with a direct type cast
custom := shared.HealthcareProcedureStatus("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `HealthcareProcedureStatusPreparation`    | PREPARATION                               |
| `HealthcareProcedureStatusInProgress`     | IN_PROGRESS                               |
| `HealthcareProcedureStatusNotDone`        | NOT_DONE                                  |
| `HealthcareProcedureStatusOnHold`         | ON_HOLD                                   |
| `HealthcareProcedureStatusStopped`        | STOPPED                                   |
| `HealthcareProcedureStatusCompleted`      | COMPLETED                                 |
| `HealthcareProcedureStatusEnteredInError` | ENTERED_IN_ERROR                          |
| `HealthcareProcedureStatusUnknown`        | UNKNOWN                                   |
| `HealthcareProcedureStatusOther`          | OTHER                                     |
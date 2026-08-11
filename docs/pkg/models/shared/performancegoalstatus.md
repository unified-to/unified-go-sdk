# PerformanceGoalStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PerformanceGoalStatusNotStarted

// Open enum: custom values can be created with a direct type cast
custom := shared.PerformanceGoalStatus("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `PerformanceGoalStatusNotStarted` | NOT_STARTED                       |
| `PerformanceGoalStatusInProgress` | IN_PROGRESS                       |
| `PerformanceGoalStatusCompleted`  | COMPLETED                         |
| `PerformanceGoalStatusClosed`     | CLOSED                            |
| `PerformanceGoalStatusOther`      | OTHER                             |
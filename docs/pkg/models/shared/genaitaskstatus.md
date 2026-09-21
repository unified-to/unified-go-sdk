# GenaiTaskStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.GenaiTaskStatusPending

// Open enum: custom values can be created with a direct type cast
custom := shared.GenaiTaskStatus("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `GenaiTaskStatusPending`   | PENDING                    |
| `GenaiTaskStatusRunning`   | RUNNING                    |
| `GenaiTaskStatusBlocked`   | BLOCKED                    |
| `GenaiTaskStatusCompleted` | COMPLETED                  |
| `GenaiTaskStatusFailed`    | FAILED                     |
| `GenaiTaskStatusCancelled` | CANCELLED                  |
| `GenaiTaskStatusExpired`   | EXPIRED                    |
| `GenaiTaskStatusOther`     | OTHER                      |
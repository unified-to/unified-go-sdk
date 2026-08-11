# PerformanceCycleStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PerformanceCycleStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := shared.PerformanceCycleStatus("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `PerformanceCycleStatusDraft`     | DRAFT                             |
| `PerformanceCycleStatusScheduled` | SCHEDULED                         |
| `PerformanceCycleStatusActive`    | ACTIVE                            |
| `PerformanceCycleStatusClosed`    | CLOSED                            |
| `PerformanceCycleStatusOther`     | OTHER                             |
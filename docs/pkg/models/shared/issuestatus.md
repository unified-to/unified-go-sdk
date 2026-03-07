# IssueStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.IssueStatusCompleted

// Open enum: custom values can be created with a direct type cast
custom := shared.IssueStatus("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `IssueStatusCompleted`  | COMPLETED               |
| `IssueStatusNew`        | NEW                     |
| `IssueStatusRoadmap`    | ROADMAP                 |
| `IssueStatusInProgress` | IN_PROGRESS             |
| `IssueStatusOnHold`     | ON_HOLD                 |
| `IssueStatusValidating` | VALIDATING              |
| `IssueStatusRejected`   | REJECTED                |
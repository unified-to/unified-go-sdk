# PerformanceReviewStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PerformanceReviewStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := shared.PerformanceReviewStatus("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `PerformanceReviewStatusDraft`        | DRAFT                                 |
| `PerformanceReviewStatusInProgress`   | IN_PROGRESS                           |
| `PerformanceReviewStatusSubmitted`    | SUBMITTED                             |
| `PerformanceReviewStatusCompleted`    | COMPLETED                             |
| `PerformanceReviewStatusShared`       | SHARED                                |
| `PerformanceReviewStatusAcknowledged` | ACKNOWLEDGED                          |
| `PerformanceReviewStatusOther`        | OTHER                                 |
# AssessmentOrderStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AssessmentOrderStatusOpen

// Open enum: custom values can be created with a direct type cast
custom := shared.AssessmentOrderStatus("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `AssessmentOrderStatusOpen`       | OPEN                              |
| `AssessmentOrderStatusInProgress` | IN_PROGRESS                       |
| `AssessmentOrderStatusCompleted`  | COMPLETED                         |
| `AssessmentOrderStatusFailed`     | FAILED                            |
| `AssessmentOrderStatusRejected`   | REJECTED                          |
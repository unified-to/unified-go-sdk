# AssessmentAttributeStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AssessmentAttributeStatusOpen

// Open enum: custom values can be created with a direct type cast
custom := shared.AssessmentAttributeStatus("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `AssessmentAttributeStatusOpen`       | OPEN                                  |
| `AssessmentAttributeStatusInProgress` | IN_PROGRESS                           |
| `AssessmentAttributeStatusCompleted`  | COMPLETED                             |
| `AssessmentAttributeStatusFailed`     | FAILED                                |
| `AssessmentAttributeStatusRejected`   | REJECTED                              |
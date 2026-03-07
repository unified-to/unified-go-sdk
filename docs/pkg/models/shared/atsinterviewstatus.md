# AtsInterviewStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AtsInterviewStatusScheduled

// Open enum: custom values can be created with a direct type cast
custom := shared.AtsInterviewStatus("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `AtsInterviewStatusScheduled`        | SCHEDULED                            |
| `AtsInterviewStatusAwaitingFeedback` | AWAITING_FEEDBACK                    |
| `AtsInterviewStatusComplete`         | COMPLETE                             |
| `AtsInterviewStatusCanceled`         | CANCELED                             |
| `AtsInterviewStatusNeedsScheduling`  | NEEDS_SCHEDULING                     |
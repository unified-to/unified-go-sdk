# AtsJobStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AtsJobStatusArchived

// Open enum: custom values can be created with a direct type cast
custom := shared.AtsJobStatus("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `AtsJobStatusArchived` | ARCHIVED               |
| `AtsJobStatusPending`  | PENDING                |
| `AtsJobStatusDraft`    | DRAFT                  |
| `AtsJobStatusOpen`     | OPEN                   |
| `AtsJobStatusClosed`   | CLOSED                 |
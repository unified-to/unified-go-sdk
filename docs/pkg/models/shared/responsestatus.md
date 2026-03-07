# ResponseStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ResponseStatusOpen

// Open enum: custom values can be created with a direct type cast
custom := shared.ResponseStatus("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `ResponseStatusOpen`       | OPEN                       |
| `ResponseStatusInProgress` | IN_PROGRESS                |
| `ResponseStatusCompleted`  | COMPLETED                  |
| `ResponseStatusFailed`     | FAILED                     |
| `ResponseStatusRejected`   | REJECTED                   |
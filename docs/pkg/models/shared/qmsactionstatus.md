# QmsActionStatus

Normalized status of a quality record

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.QmsActionStatusOpen

// Open enum: custom values can be created with a direct type cast
custom := shared.QmsActionStatus("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `QmsActionStatusOpen`       | OPEN                        |
| `QmsActionStatusInProgress` | IN_PROGRESS                 |
| `QmsActionStatusInReview`   | IN_REVIEW                   |
| `QmsActionStatusClosed`     | CLOSED                      |
| `QmsActionStatusCancelled`  | CANCELLED                   |
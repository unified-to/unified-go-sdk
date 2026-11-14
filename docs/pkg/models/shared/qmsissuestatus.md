# QmsIssueStatus

Normalized status of a quality record

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.QmsIssueStatusOpen

// Open enum: custom values can be created with a direct type cast
custom := shared.QmsIssueStatus("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `QmsIssueStatusOpen`       | OPEN                       |
| `QmsIssueStatusInProgress` | IN_PROGRESS                |
| `QmsIssueStatusInReview`   | IN_REVIEW                  |
| `QmsIssueStatusClosed`     | CLOSED                     |
| `QmsIssueStatusCancelled`  | CANCELLED                  |
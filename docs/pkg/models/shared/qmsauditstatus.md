# QmsAuditStatus

Normalized status of a quality record

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.QmsAuditStatusOpen

// Open enum: custom values can be created with a direct type cast
custom := shared.QmsAuditStatus("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `QmsAuditStatusOpen`       | OPEN                       |
| `QmsAuditStatusInProgress` | IN_PROGRESS                |
| `QmsAuditStatusInReview`   | IN_REVIEW                  |
| `QmsAuditStatusClosed`     | CLOSED                     |
| `QmsAuditStatusCancelled`  | CANCELLED                  |
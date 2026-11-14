# QmsInspectionStatus

Normalized status of a quality record

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.QmsInspectionStatusOpen

// Open enum: custom values can be created with a direct type cast
custom := shared.QmsInspectionStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `QmsInspectionStatusOpen`       | OPEN                            |
| `QmsInspectionStatusInProgress` | IN_PROGRESS                     |
| `QmsInspectionStatusInReview`   | IN_REVIEW                       |
| `QmsInspectionStatusClosed`     | CLOSED                          |
| `QmsInspectionStatusCancelled`  | CANCELLED                       |
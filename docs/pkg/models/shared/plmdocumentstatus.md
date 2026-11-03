# PlmDocumentStatus

Normalized lifecycle status of a part or document revision

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PlmDocumentStatusInWork

// Open enum: custom values can be created with a direct type cast
custom := shared.PlmDocumentStatus("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `PlmDocumentStatusInWork`   | IN_WORK                     |
| `PlmDocumentStatusInReview` | IN_REVIEW                   |
| `PlmDocumentStatusReleased` | RELEASED                    |
| `PlmDocumentStatusObsolete` | OBSOLETE                    |
| `PlmDocumentStatusOther`    | OTHER                       |
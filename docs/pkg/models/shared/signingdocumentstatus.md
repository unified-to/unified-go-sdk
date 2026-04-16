# SigningDocumentStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.SigningDocumentStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := shared.SigningDocumentStatus("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `SigningDocumentStatusDraft`      | DRAFT                             |
| `SigningDocumentStatusSent`       | SENT                              |
| `SigningDocumentStatusDelivered`  | DELIVERED                         |
| `SigningDocumentStatusInProgress` | IN_PROGRESS                       |
| `SigningDocumentStatusCompleted`  | COMPLETED                         |
| `SigningDocumentStatusDeclined`   | DECLINED                          |
| `SigningDocumentStatusVoided`     | VOIDED                            |
| `SigningDocumentStatusExpired`    | EXPIRED                           |
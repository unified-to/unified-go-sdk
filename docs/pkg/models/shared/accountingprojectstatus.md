# AccountingProjectStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AccountingProjectStatusNotStarted

// Open enum: custom values can be created with a direct type cast
custom := shared.AccountingProjectStatus("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `AccountingProjectStatusNotStarted` | NOT_STARTED                         |
| `AccountingProjectStatusInProgress` | IN_PROGRESS                         |
| `AccountingProjectStatusOnHold`     | ON_HOLD                             |
| `AccountingProjectStatusCompleted`  | COMPLETED                           |
| `AccountingProjectStatusCancelled`  | CANCELLED                           |
| `AccountingProjectStatusClosed`     | CLOSED                              |
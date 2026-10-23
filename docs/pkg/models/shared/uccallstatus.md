# UcCallStatus

The current or final status of the call

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.UcCallStatusQueued

// Open enum: custom values can be created with a direct type cast
custom := shared.UcCallStatus("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `UcCallStatusQueued`     | QUEUED                   |
| `UcCallStatusRinging`    | RINGING                  |
| `UcCallStatusInProgress` | IN_PROGRESS              |
| `UcCallStatusCompleted`  | COMPLETED                |
| `UcCallStatusBusy`       | BUSY                     |
| `UcCallStatusNoAnswer`   | NO_ANSWER                |
| `UcCallStatusFailed`     | FAILED                   |
| `UcCallStatusCanceled`   | CANCELED                 |
| `UcCallStatusOther`      | OTHER                    |
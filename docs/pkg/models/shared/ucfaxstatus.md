# UcFaxStatus

The current or final status of the fax

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.UcFaxStatusQueued

// Open enum: custom values can be created with a direct type cast
custom := shared.UcFaxStatus("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `UcFaxStatusQueued`    | QUEUED                 |
| `UcFaxStatusSending`   | SENDING                |
| `UcFaxStatusDelivered` | DELIVERED              |
| `UcFaxStatusReceiving` | RECEIVING              |
| `UcFaxStatusReceived`  | RECEIVED               |
| `UcFaxStatusFailed`    | FAILED                 |
| `UcFaxStatusCanceled`  | CANCELED               |
| `UcFaxStatusOther`     | OTHER                  |
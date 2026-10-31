# CommerceReturnStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.CommerceReturnStatusRequested

// Open enum: custom values can be created with a direct type cast
custom := shared.CommerceReturnStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `CommerceReturnStatusRequested` | REQUESTED                       |
| `CommerceReturnStatusOpen`      | OPEN                            |
| `CommerceReturnStatusApproved`  | APPROVED                        |
| `CommerceReturnStatusRejected`  | REJECTED                        |
| `CommerceReturnStatusClosed`    | CLOSED                          |
| `CommerceReturnStatusCancelled` | CANCELLED                       |
| `CommerceReturnStatusOther`     | OTHER                           |
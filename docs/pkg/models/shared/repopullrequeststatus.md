# RepoPullrequestStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.RepoPullrequestStatusPending

// Open enum: custom values can be created with a direct type cast
custom := shared.RepoPullrequestStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `RepoPullrequestStatusPending`  | PENDING                         |
| `RepoPullrequestStatusApproved` | APPROVED                        |
| `RepoPullrequestStatusRejected` | REJECTED                        |
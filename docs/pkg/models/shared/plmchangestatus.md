# PlmChangeStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PlmChangeStatusOpen

// Open enum: custom values can be created with a direct type cast
custom := shared.PlmChangeStatus("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `PlmChangeStatusOpen`        | OPEN                         |
| `PlmChangeStatusInReview`    | IN_REVIEW                    |
| `PlmChangeStatusApproved`    | APPROVED                     |
| `PlmChangeStatusImplemented` | IMPLEMENTED                  |
| `PlmChangeStatusRejected`    | REJECTED                     |
| `PlmChangeStatusCancelled`   | CANCELLED                    |
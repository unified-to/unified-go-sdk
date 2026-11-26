# DmsRepairorderJobStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.DmsRepairorderJobStatusPending

// Open enum: custom values can be created with a direct type cast
custom := shared.DmsRepairorderJobStatus("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `DmsRepairorderJobStatusPending`    | PENDING                             |
| `DmsRepairorderJobStatusInProgress` | IN_PROGRESS                         |
| `DmsRepairorderJobStatusOnHold`     | ON_HOLD                             |
| `DmsRepairorderJobStatusCompleted`  | COMPLETED                           |
| `DmsRepairorderJobStatusDeclined`   | DECLINED                            |
| `DmsRepairorderJobStatusCancelled`  | CANCELLED                           |
| `DmsRepairorderJobStatusOther`      | OTHER                               |
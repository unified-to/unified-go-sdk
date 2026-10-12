# ComputeSnapshotStatus

Normalized snapshot status

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ComputeSnapshotStatusCreating

// Open enum: custom values can be created with a direct type cast
custom := shared.ComputeSnapshotStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ComputeSnapshotStatusCreating` | CREATING                        |
| `ComputeSnapshotStatusReady`    | READY                           |
| `ComputeSnapshotStatusFailed`   | FAILED                          |
| `ComputeSnapshotStatusDeleting` | DELETING                        |
| `ComputeSnapshotStatusOther`    | OTHER                           |
# ComputeSnapshotType

Whether memory state is included

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ComputeSnapshotTypeMemory

// Open enum: custom values can be created with a direct type cast
custom := shared.ComputeSnapshotType("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `ComputeSnapshotTypeMemory` | MEMORY                      |
| `ComputeSnapshotTypeDisk`   | DISK                        |
| `ComputeSnapshotTypeOther`  | OTHER                       |
# ComputeVolumeStatus

Normalized volume status

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ComputeVolumeStatusCreating

// Open enum: custom values can be created with a direct type cast
custom := shared.ComputeVolumeStatus("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `ComputeVolumeStatusCreating`  | CREATING                       |
| `ComputeVolumeStatusAvailable` | AVAILABLE                      |
| `ComputeVolumeStatusInUse`     | IN_USE                         |
| `ComputeVolumeStatusDeleting`  | DELETING                       |
| `ComputeVolumeStatusError`     | ERROR                          |
| `ComputeVolumeStatusOther`     | OTHER                          |
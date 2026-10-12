# ComputeVolumeType

The kind of volume

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ComputeVolumeTypeStandard

// Open enum: custom values can be created with a direct type cast
custom := shared.ComputeVolumeType("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `ComputeVolumeTypeStandard`        | STANDARD                           |
| `ComputeVolumeTypeHighPerformance` | HIGH_PERFORMANCE                   |
| `ComputeVolumeTypeShared`          | SHARED                             |
| `ComputeVolumeTypeLocal`           | LOCAL                              |
| `ComputeVolumeTypeOther`           | OTHER                              |
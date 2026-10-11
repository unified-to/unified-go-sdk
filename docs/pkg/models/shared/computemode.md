# ComputeMode

How the segment membership is computed (REALTIME or BATCH)

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ComputeModeRealtime

// Open enum: custom values can be created with a direct type cast
custom := shared.ComputeMode("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `ComputeModeRealtime` | REALTIME              |
| `ComputeModeBatch`    | BATCH                 |
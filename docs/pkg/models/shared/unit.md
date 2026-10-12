# Unit

Unit of quantity

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.UnitGpuHour

// Open enum: custom values can be created with a direct type cast
custom := shared.Unit("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `UnitGpuHour`    | GPU_HOUR         |
| `UnitVcpuSecond` | VCPU_SECOND      |
| `UnitGbSecond`   | GB_SECOND        |
| `UnitGbMonth`    | GB_MONTH         |
| `UnitByte`       | BYTE             |
| `UnitSecond`     | SECOND           |
| `UnitOther`      | OTHER            |
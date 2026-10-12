# PriceUnit

Unit that price_amount is quoted in

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PriceUnitPerGpuHour

// Open enum: custom values can be created with a direct type cast
custom := shared.PriceUnit("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `PriceUnitPerGpuHour`      | PER_GPU_HOUR               |
| `PriceUnitPerInstanceHour` | PER_INSTANCE_HOUR          |
| `PriceUnitOther`           | OTHER                      |
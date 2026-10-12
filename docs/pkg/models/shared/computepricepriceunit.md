# ComputePricePriceUnit

Unit that price_amount is quoted in

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ComputePricePriceUnitPerGpuHour

// Open enum: custom values can be created with a direct type cast
custom := shared.ComputePricePriceUnit("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `ComputePricePriceUnitPerGpuHour`      | PER_GPU_HOUR                           |
| `ComputePricePriceUnitPerInstanceHour` | PER_INSTANCE_HOUR                      |
| `ComputePricePriceUnitOther`           | OTHER                                  |
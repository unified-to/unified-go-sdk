# ComputePriceMarket

Pricing market

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ComputePriceMarketOnDemand

// Open enum: custom values can be created with a direct type cast
custom := shared.ComputePriceMarket("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `ComputePriceMarketOnDemand`  | ON_DEMAND                     |
| `ComputePriceMarketSpot`      | SPOT                          |
| `ComputePriceMarketReserved`  | RESERVED                      |
| `ComputePriceMarketCommunity` | COMMUNITY                     |
| `ComputePriceMarketOther`     | OTHER                         |
# Market

Pricing market the instance runs in

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.MarketOnDemand

// Open enum: custom values can be created with a direct type cast
custom := shared.Market("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `MarketOnDemand`  | ON_DEMAND         |
| `MarketSpot`      | SPOT              |
| `MarketReserved`  | RESERVED          |
| `MarketCommunity` | COMMUNITY         |
| `MarketOther`     | OTHER             |
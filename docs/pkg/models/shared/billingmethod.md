# BillingMethod

how the matter is billed

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.BillingMethodHourly

// Open enum: custom values can be created with a direct type cast
custom := shared.BillingMethod("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `BillingMethodHourly`      | HOURLY                     |
| `BillingMethodFlatFee`     | FLAT_FEE                   |
| `BillingMethodContingency` | CONTINGENCY                |
| `BillingMethodOther`       | OTHER                      |
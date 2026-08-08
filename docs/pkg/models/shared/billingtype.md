# BillingType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.BillingTypeFixedPrice

// Open enum: custom values can be created with a direct type cast
custom := shared.BillingType("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `BillingTypeFixedPrice`       | FIXED_PRICE                   |
| `BillingTypeTimeAndMaterials` | TIME_AND_MATERIALS            |
| `BillingTypeMilestone`        | MILESTONE                     |
| `BillingTypeNonBillable`      | NON_BILLABLE                  |
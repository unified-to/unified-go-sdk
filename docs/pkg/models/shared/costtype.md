# CostType

How the line item is priced (Google Ad Manager costType)

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.CostTypeCpm

// Open enum: custom values can be created with a direct type cast
custom := shared.CostType("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `CostTypeCpm`         | CPM                   |
| `CostTypeViewableCpm` | VIEWABLE_CPM          |
| `CostTypeCpc`         | CPC                   |
| `CostTypeCpd`         | CPD                   |
| `CostTypeCpa`         | CPA                   |
| `CostTypeCpcv`        | CPCV                  |
| `CostTypeOther`       | OTHER                 |
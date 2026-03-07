# BudgetAllocationType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.BudgetAllocationTypeUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.BudgetAllocationType("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `BudgetAllocationTypeUnspecified` | UNSPECIFIED                       |
| `BudgetAllocationTypeAutomatic`   | AUTOMATIC                         |
| `BudgetAllocationTypeFixed`       | FIXED                             |
| `BudgetAllocationTypeUnlimited`   | UNLIMITED                         |
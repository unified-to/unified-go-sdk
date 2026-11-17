# OverrideType

replaces the calculated withholding with a flat amount or percentage

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.OverrideTypeNone

// Open enum: custom values can be created with a direct type cast
custom := shared.OverrideType("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `OverrideTypeNone`        | NONE                      |
| `OverrideTypeFixedAmount` | FIXED_AMOUNT              |
| `OverrideTypePercent`     | PERCENT                   |
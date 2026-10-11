# SplitType

how net pay is split into this account when an employee has several accounts

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.SplitTypePercentage

// Open enum: custom values can be created with a direct type cast
custom := shared.SplitType("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `SplitTypePercentage` | PERCENTAGE            |
| `SplitTypeFixed`      | FIXED                 |
| `SplitTypeRemainder`  | REMAINDER             |
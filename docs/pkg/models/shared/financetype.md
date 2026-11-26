# FinanceType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.FinanceTypeCash

// Open enum: custom values can be created with a direct type cast
custom := shared.FinanceType("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `FinanceTypeCash`    | CASH                 |
| `FinanceTypeFinance` | FINANCE              |
| `FinanceTypeLease`   | LEASE                |
| `FinanceTypeBalloon` | BALLOON              |
| `FinanceTypeOther`   | OTHER                |
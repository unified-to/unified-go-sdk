# AccountingExpenseType

USER = employee spend; COMPANY = books/purchase expense

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AccountingExpenseTypeUser

// Open enum: custom values can be created with a direct type cast
custom := shared.AccountingExpenseType("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `AccountingExpenseTypeUser`    | USER                           |
| `AccountingExpenseTypeCompany` | COMPANY                        |
# AccountingExpenseStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AccountingExpenseStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := shared.AccountingExpenseStatus("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `AccountingExpenseStatusDraft`     | DRAFT                              |
| `AccountingExpenseStatusSubmitted` | SUBMITTED                          |
| `AccountingExpenseStatusPending`   | PENDING                            |
| `AccountingExpenseStatusApproved`  | APPROVED                           |
| `AccountingExpenseStatusRejected`  | REJECTED                           |
| `AccountingExpenseStatusPaid`      | PAID                               |
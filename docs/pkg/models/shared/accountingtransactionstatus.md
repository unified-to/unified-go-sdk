# AccountingTransactionStatus

Posting state of the transaction

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AccountingTransactionStatusPending

// Open enum: custom values can be created with a direct type cast
custom := shared.AccountingTransactionStatus("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `AccountingTransactionStatusPending` | PENDING                              |
| `AccountingTransactionStatusPosted`  | POSTED                               |
| `AccountingTransactionStatusVoid`    | VOID                                 |
| `AccountingTransactionStatusOther`   | OTHER                                |
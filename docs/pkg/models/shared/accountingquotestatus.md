# AccountingQuoteStatus

where EXPIRED is derived (expires_at < now), the source status is preserved in raw

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AccountingQuoteStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := shared.AccountingQuoteStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `AccountingQuoteStatusDraft`    | DRAFT                           |
| `AccountingQuoteStatusSent`     | SENT                            |
| `AccountingQuoteStatusAccepted` | ACCEPTED                        |
| `AccountingQuoteStatusDeclined` | DECLINED                        |
| `AccountingQuoteStatusInvoiced` | INVOICED                        |
| `AccountingQuoteStatusExpired`  | EXPIRED                         |
| `AccountingQuoteStatusVoided`   | VOIDED                          |
| `AccountingQuoteStatusDeleted`  | DELETED                         |
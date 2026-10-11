# AccountingJournalStatus

Posting state of the journal entry

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AccountingJournalStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := shared.AccountingJournalStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `AccountingJournalStatusDraft`  | DRAFT                           |
| `AccountingJournalStatusPosted` | POSTED                          |
| `AccountingJournalStatusVoid`   | VOID                            |
| `AccountingJournalStatusOther`  | OTHER                           |
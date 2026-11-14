# QmsIssueSource

Where a quality issue was found

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.QmsIssueSourceInternal

// Open enum: custom values can be created with a direct type cast
custom := shared.QmsIssueSource("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `QmsIssueSourceInternal` | INTERNAL                 |
| `QmsIssueSourceSupplier` | SUPPLIER                 |
| `QmsIssueSourceCustomer` | CUSTOMER                 |
| `QmsIssueSourceAudit`    | AUDIT                    |
| `QmsIssueSourceOther`    | OTHER                    |
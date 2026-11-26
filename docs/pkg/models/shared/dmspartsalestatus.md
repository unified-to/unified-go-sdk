# DmsPartsaleStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.DmsPartsaleStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := shared.DmsPartsaleStatus("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `DmsPartsaleStatusDraft`    | DRAFT                       |
| `DmsPartsaleStatusOpen`     | OPEN                        |
| `DmsPartsaleStatusInvoiced` | INVOICED                    |
| `DmsPartsaleStatusShipped`  | SHIPPED                     |
| `DmsPartsaleStatusClosed`   | CLOSED                      |
| `DmsPartsaleStatusVoided`   | VOIDED                      |
| `DmsPartsaleStatusOther`    | OTHER                       |
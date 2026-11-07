# UcFaxType

Whether the fax was received (INBOUND) or sent (OUTBOUND)

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.UcFaxTypeInbound

// Open enum: custom values can be created with a direct type cast
custom := shared.UcFaxType("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `UcFaxTypeInbound`  | INBOUND             |
| `UcFaxTypeOutbound` | OUTBOUND            |
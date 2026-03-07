# TaxesPaidBy

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.TaxesPaidBySender

// Open enum: custom values can be created with a direct type cast
custom := shared.TaxesPaidBy("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `TaxesPaidBySender`     | SENDER                  |
| `TaxesPaidByRecipient`  | RECIPIENT               |
| `TaxesPaidByThirdParty` | THIRD_PARTY             |
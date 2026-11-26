# DmsOpcodePayType

Default pay type

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.DmsOpcodePayTypeCustomer

// Open enum: custom values can be created with a direct type cast
custom := shared.DmsOpcodePayType("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `DmsOpcodePayTypeCustomer`        | CUSTOMER                          |
| `DmsOpcodePayTypeWarranty`        | WARRANTY                          |
| `DmsOpcodePayTypeInternal`        | INTERNAL                          |
| `DmsOpcodePayTypeServiceContract` | SERVICE_CONTRACT                  |
| `DmsOpcodePayTypeOther`           | OTHER                             |
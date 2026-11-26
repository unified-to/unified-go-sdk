# DmsPartlinePayType

Who pays for this line

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.DmsPartlinePayTypeCustomer

// Open enum: custom values can be created with a direct type cast
custom := shared.DmsPartlinePayType("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `DmsPartlinePayTypeCustomer`        | CUSTOMER                            |
| `DmsPartlinePayTypeWarranty`        | WARRANTY                            |
| `DmsPartlinePayTypeInternal`        | INTERNAL                            |
| `DmsPartlinePayTypeServiceContract` | SERVICE_CONTRACT                    |
| `DmsPartlinePayTypeOther`           | OTHER                               |
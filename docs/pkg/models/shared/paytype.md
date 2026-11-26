# PayType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PayTypeCustomer

// Open enum: custom values can be created with a direct type cast
custom := shared.PayType("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `PayTypeCustomer`        | CUSTOMER                 |
| `PayTypeWarranty`        | WARRANTY                 |
| `PayTypeInternal`        | INTERNAL                 |
| `PayTypeServiceContract` | SERVICE_CONTRACT         |
| `PayTypeOther`           | OTHER                    |
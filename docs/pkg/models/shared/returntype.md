# ReturnType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.ReturnTypeCustomer

// Open enum: custom values can be created with a direct type cast
custom := shared.ReturnType("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `ReturnTypeCustomer`  | CUSTOMER              |
| `ReturnTypeVendor`    | VENDOR                |
| `ReturnTypeWarranty`  | WARRANTY              |
| `ReturnTypeDefective` | DEFECTIVE             |
| `ReturnTypeOther`     | OTHER                 |
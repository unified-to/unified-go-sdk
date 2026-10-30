# PropertyLineitemType

base price, fee, tax, discount or deposit

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PropertyLineitemTypeBase

// Open enum: custom values can be created with a direct type cast
custom := shared.PropertyLineitemType("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `PropertyLineitemTypeBase`     | BASE                           |
| `PropertyLineitemTypeFee`      | FEE                            |
| `PropertyLineitemTypeTax`      | TAX                            |
| `PropertyLineitemTypeDiscount` | DISCOUNT                       |
| `PropertyLineitemTypeDeposit`  | DEPOSIT                        |
| `PropertyLineitemTypeOther`    | OTHER                          |
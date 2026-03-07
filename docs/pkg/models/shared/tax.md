# Tax

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.TaxPreTax

// Open enum: custom values can be created with a direct type cast
custom := shared.Tax("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `TaxPreTax`     | PRE_TAX         |
| `TaxPostTax`    | POST_TAX        |
| `TaxTaxable`    | TAXABLE         |
| `TaxNonTaxable` | NON_TAXABLE     |
| `TaxTax`        | TAX             |
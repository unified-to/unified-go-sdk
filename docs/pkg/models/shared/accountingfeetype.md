# AccountingFeeType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AccountingFeeTypeTax

// Open enum: custom values can be created with a direct type cast
custom := shared.AccountingFeeType("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `AccountingFeeTypeTax`       | TAX                          |
| `AccountingFeeTypeDiscount`  | DISCOUNT                     |
| `AccountingFeeTypePromotion` | PROMOTION                    |
| `AccountingFeeTypeShipping`  | SHIPPING                     |
| `AccountingFeeTypeGiftWrap`  | GIFT_WRAP                    |
| `AccountingFeeTypeCod`       | COD                          |
| `AccountingFeeTypeSurcharge` | SURCHARGE                    |
| `AccountingFeeTypeOther`     | OTHER                        |
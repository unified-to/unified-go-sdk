# PaymentPayoutLineitemObjectType

What this payout line settles

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PaymentPayoutLineitemObjectTypePayment

// Open enum: custom values can be created with a direct type cast
custom := shared.PaymentPayoutLineitemObjectType("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `PaymentPayoutLineitemObjectTypePayment`    | PAYMENT                                     |
| `PaymentPayoutLineitemObjectTypeRefund`     | REFUND                                      |
| `PaymentPayoutLineitemObjectTypeFee`        | FEE                                         |
| `PaymentPayoutLineitemObjectTypeAdjustment` | ADJUSTMENT                                  |
| `PaymentPayoutLineitemObjectTypeTransfer`   | TRANSFER                                    |
| `PaymentPayoutLineitemObjectTypeOther`      | OTHER                                       |
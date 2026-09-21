# PaymentRefundTenderType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PaymentRefundTenderTypeCard

// Open enum: custom values can be created with a direct type cast
custom := shared.PaymentRefundTenderType("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `PaymentRefundTenderTypeCard`           | CARD                                    |
| `PaymentRefundTenderTypeCash`           | CASH                                    |
| `PaymentRefundTenderTypeGiftCard`       | GIFT_CARD                               |
| `PaymentRefundTenderTypeBankTransfer`   | BANK_TRANSFER                           |
| `PaymentRefundTenderTypeWallet`         | WALLET                                  |
| `PaymentRefundTenderTypeCheck`          | CHECK                                   |
| `PaymentRefundTenderTypeStoreCredit`    | STORE_CREDIT                            |
| `PaymentRefundTenderTypeBuyNowPayLater` | BUY_NOW_PAY_LATER                       |
| `PaymentRefundTenderTypeExternal`       | EXTERNAL                                |
| `PaymentRefundTenderTypeOther`          | OTHER                                   |
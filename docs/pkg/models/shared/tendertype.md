# TenderType

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.TenderTypeCard

// Open enum: custom values can be created with a direct type cast
custom := shared.TenderType("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `TenderTypeCard`           | CARD                       |
| `TenderTypeCash`           | CASH                       |
| `TenderTypeGiftCard`       | GIFT_CARD                  |
| `TenderTypeBankTransfer`   | BANK_TRANSFER              |
| `TenderTypeWallet`         | WALLET                     |
| `TenderTypeCheck`          | CHECK                      |
| `TenderTypeStoreCredit`    | STORE_CREDIT               |
| `TenderTypeBuyNowPayLater` | BUY_NOW_PAY_LATER          |
| `TenderTypeExternal`       | EXTERNAL                   |
| `TenderTypeOther`          | OTHER                      |
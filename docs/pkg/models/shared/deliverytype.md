# DeliveryType

Ad server line item type / delivery class (Google Ad Manager lineItemType)

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.DeliveryTypeStandard

// Open enum: custom values can be created with a direct type cast
custom := shared.DeliveryType("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `DeliveryTypeStandard`      | STANDARD                    |
| `DeliveryTypeSponsorship`   | SPONSORSHIP                 |
| `DeliveryTypeNetwork`       | NETWORK                     |
| `DeliveryTypeBulk`          | BULK                        |
| `DeliveryTypePricePriority` | PRICE_PRIORITY              |
| `DeliveryTypeHouse`         | HOUSE                       |
| `DeliveryTypePreferredDeal` | PREFERRED_DEAL              |
| `DeliveryTypeAdExchange`    | AD_EXCHANGE                 |
| `DeliveryTypeClickTracking` | CLICK_TRACKING              |
| `DeliveryTypeOther`         | OTHER                       |
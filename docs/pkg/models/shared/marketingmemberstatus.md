# MarketingMemberStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.MarketingMemberStatusSubscribed

// Open enum: custom values can be created with a direct type cast
custom := shared.MarketingMemberStatus("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `MarketingMemberStatusSubscribed`    | SUBSCRIBED                           |
| `MarketingMemberStatusUnsubscribed`  | UNSUBSCRIBED                         |
| `MarketingMemberStatusCleaned`       | CLEANED                              |
| `MarketingMemberStatusPending`       | PENDING                              |
| `MarketingMemberStatusTransactional` | TRANSACTIONAL                        |
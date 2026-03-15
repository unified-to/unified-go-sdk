# BillingEvent

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.BillingEventImpressions

// Open enum: custom values can be created with a direct type cast
custom := shared.BillingEvent("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `BillingEventImpressions`    | IMPRESSIONS                  |
| `BillingEventLinkClicks`     | LINK_CLICKS                  |
| `BillingEventVideoViews`     | VIDEO_VIEWS                  |
| `BillingEventAppInstalls`    | APP_INSTALLS                 |
| `BillingEventEngagement`     | ENGAGEMENT                   |
| `BillingEventPageLikes`      | PAGE_LIKES                   |
| `BillingEventMessages`       | MESSAGES                     |
| `BillingEventPostEngagement` | POST_ENGAGEMENT              |
| `BillingEventPurchase`       | PURCHASE                     |
| `BillingEventNone`           | NONE                         |
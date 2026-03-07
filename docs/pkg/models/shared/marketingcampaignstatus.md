# MarketingCampaignStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.MarketingCampaignStatusDraft

// Open enum: custom values can be created with a direct type cast
custom := shared.MarketingCampaignStatus("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `MarketingCampaignStatusDraft`     | DRAFT                              |
| `MarketingCampaignStatusScheduled` | SCHEDULED                          |
| `MarketingCampaignStatusSending`   | SENDING                            |
| `MarketingCampaignStatusSent`      | SENT                               |
| `MarketingCampaignStatusCancelled` | CANCELLED                          |
| `MarketingCampaignStatusPaused`    | PAUSED                             |
| `MarketingCampaignStatusArchived`  | ARCHIVED                           |
| `MarketingCampaignStatusOther`     | OTHER                              |
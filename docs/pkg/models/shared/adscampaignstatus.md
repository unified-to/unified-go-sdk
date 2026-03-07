# AdsCampaignStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AdsCampaignStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AdsCampaignStatus("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `AdsCampaignStatusUnspecified`          | UNSPECIFIED                             |
| `AdsCampaignStatusActive`               | ACTIVE                                  |
| `AdsCampaignStatusPaused`               | PAUSED                                  |
| `AdsCampaignStatusArchived`             | ARCHIVED                                |
| `AdsCampaignStatusDraft`                | DRAFT                                   |
| `AdsCampaignStatusScheduledForDeletion` | SCHEDULED_FOR_DELETION                  |
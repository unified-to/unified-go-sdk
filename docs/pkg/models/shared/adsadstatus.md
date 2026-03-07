# AdsAdStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AdsAdStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AdsAdStatus("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `AdsAdStatusUnspecified`          | UNSPECIFIED                       |
| `AdsAdStatusActive`               | ACTIVE                            |
| `AdsAdStatusPaused`               | PAUSED                            |
| `AdsAdStatusArchived`             | ARCHIVED                          |
| `AdsAdStatusDraft`                | DRAFT                             |
| `AdsAdStatusScheduledForDeletion` | SCHEDULED_FOR_DELETION            |
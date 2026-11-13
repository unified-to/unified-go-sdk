# AdsInventoryStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AdsInventoryStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AdsInventoryStatus("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `AdsInventoryStatusUnspecified`          | UNSPECIFIED                              |
| `AdsInventoryStatusActive`               | ACTIVE                                   |
| `AdsInventoryStatusPaused`               | PAUSED                                   |
| `AdsInventoryStatusArchived`             | ARCHIVED                                 |
| `AdsInventoryStatusDraft`                | DRAFT                                    |
| `AdsInventoryStatusScheduledForDeletion` | SCHEDULED_FOR_DELETION                   |
| `AdsInventoryStatusProcessing`           | PROCESSING                               |
| `AdsInventoryStatusProcessingFailed`     | PROCESSING_FAILED                        |
# AdsGroupStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AdsGroupStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AdsGroupStatus("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `AdsGroupStatusUnspecified`          | UNSPECIFIED                          |
| `AdsGroupStatusActive`               | ACTIVE                               |
| `AdsGroupStatusPaused`               | PAUSED                               |
| `AdsGroupStatusArchived`             | ARCHIVED                             |
| `AdsGroupStatusDraft`                | DRAFT                                |
| `AdsGroupStatusScheduledForDeletion` | SCHEDULED_FOR_DELETION               |
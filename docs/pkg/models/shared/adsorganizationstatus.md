# AdsOrganizationStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AdsOrganizationStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AdsOrganizationStatus("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `AdsOrganizationStatusUnspecified`          | UNSPECIFIED                                 |
| `AdsOrganizationStatusActive`               | ACTIVE                                      |
| `AdsOrganizationStatusPaused`               | PAUSED                                      |
| `AdsOrganizationStatusArchived`             | ARCHIVED                                    |
| `AdsOrganizationStatusDraft`                | DRAFT                                       |
| `AdsOrganizationStatusScheduledForDeletion` | SCHEDULED_FOR_DELETION                      |
| `AdsOrganizationStatusProcessing`           | PROCESSING                                  |
| `AdsOrganizationStatusProcessingFailed`     | PROCESSING_FAILED                           |
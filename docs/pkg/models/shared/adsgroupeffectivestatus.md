# AdsGroupEffectiveStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.AdsGroupEffectiveStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.AdsGroupEffectiveStatus("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `AdsGroupEffectiveStatusUnspecified`   | UNSPECIFIED                            |
| `AdsGroupEffectiveStatusServing`       | SERVING                                |
| `AdsGroupEffectiveStatusLimited`       | LIMITED                                |
| `AdsGroupEffectiveStatusLearning`      | LEARNING                               |
| `AdsGroupEffectiveStatusPaused`        | PAUSED                                 |
| `AdsGroupEffectiveStatusPending`       | PENDING                                |
| `AdsGroupEffectiveStatusEnded`         | ENDED                                  |
| `AdsGroupEffectiveStatusMisconfigured` | MISCONFIGURED                          |
| `AdsGroupEffectiveStatusNotEligible`   | NOT_ELIGIBLE                           |
| `AdsGroupEffectiveStatusArchived`      | ARCHIVED                               |
| `AdsGroupEffectiveStatusRemoved`       | REMOVED                                |
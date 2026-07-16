# EffectiveStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.EffectiveStatusUnspecified

// Open enum: custom values can be created with a direct type cast
custom := shared.EffectiveStatus("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `EffectiveStatusUnspecified`   | UNSPECIFIED                    |
| `EffectiveStatusServing`       | SERVING                        |
| `EffectiveStatusLimited`       | LIMITED                        |
| `EffectiveStatusLearning`      | LEARNING                       |
| `EffectiveStatusPaused`        | PAUSED                         |
| `EffectiveStatusPending`       | PENDING                        |
| `EffectiveStatusEnded`         | ENDED                          |
| `EffectiveStatusMisconfigured` | MISCONFIGURED                  |
| `EffectiveStatusNotEligible`   | NOT_ELIGIBLE                   |
| `EffectiveStatusArchived`      | ARCHIVED                       |
| `EffectiveStatusRemoved`       | REMOVED                        |
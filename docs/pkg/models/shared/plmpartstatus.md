# PlmPartStatus

Normalized lifecycle status

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PlmPartStatusInWork

// Open enum: custom values can be created with a direct type cast
custom := shared.PlmPartStatus("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `PlmPartStatusInWork`   | IN_WORK                 |
| `PlmPartStatusInReview` | IN_REVIEW               |
| `PlmPartStatusReleased` | RELEASED                |
| `PlmPartStatusObsolete` | OBSOLETE                |
| `PlmPartStatusOther`    | OTHER                   |
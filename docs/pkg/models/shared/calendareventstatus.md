# CalendarEventStatus

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.CalendarEventStatusCanceled

// Open enum: custom values can be created with a direct type cast
custom := shared.CalendarEventStatus("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `CalendarEventStatusCanceled`  | CANCELED                       |
| `CalendarEventStatusConfirmed` | CONFIRMED                      |
| `CalendarEventStatusTentative` | TENTATIVE                      |
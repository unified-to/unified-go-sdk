# Format

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.FormatText

// Open enum: custom values can be created with a direct type cast
custom := shared.Format("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `FormatText`           | TEXT                   |
| `FormatNumber`         | NUMBER                 |
| `FormatDate`           | DATE                   |
| `FormatBoolean`        | BOOLEAN                |
| `FormatFile`           | FILE                   |
| `FormatTextarea`       | TEXTAREA               |
| `FormatSingleSelect`   | SINGLE_SELECT          |
| `FormatMultipleSelect` | MULTIPLE_SELECT        |
| `FormatMeasurement`    | MEASUREMENT            |
| `FormatPrice`          | PRICE                  |
| `FormatYesNo`          | YES_NO                 |
| `FormatCurrency`       | CURRENCY               |
| `FormatURL`            | URL                    |
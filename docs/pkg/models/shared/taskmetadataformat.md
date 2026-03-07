# TaskMetadataFormat

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.TaskMetadataFormatText

// Open enum: custom values can be created with a direct type cast
custom := shared.TaskMetadataFormat("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `TaskMetadataFormatText`           | TEXT                               |
| `TaskMetadataFormatNumber`         | NUMBER                             |
| `TaskMetadataFormatDate`           | DATE                               |
| `TaskMetadataFormatBoolean`        | BOOLEAN                            |
| `TaskMetadataFormatFile`           | FILE                               |
| `TaskMetadataFormatTextarea`       | TEXTAREA                           |
| `TaskMetadataFormatSingleSelect`   | SINGLE_SELECT                      |
| `TaskMetadataFormatMultipleSelect` | MULTIPLE_SELECT                    |
| `TaskMetadataFormatMeasurement`    | MEASUREMENT                        |
| `TaskMetadataFormatPrice`          | PRICE                              |
| `TaskMetadataFormatYesNo`          | YES_NO                             |
| `TaskMetadataFormatCurrency`       | CURRENCY                           |
| `TaskMetadataFormatURL`            | URL                                |
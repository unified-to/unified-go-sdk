# PlmMetadataFormat

The data type of the custom field, which determines how `value` is formatted.

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PlmMetadataFormatText

// Open enum: custom values can be created with a direct type cast
custom := shared.PlmMetadataFormat("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `PlmMetadataFormatText`           | TEXT                              |
| `PlmMetadataFormatNumber`         | NUMBER                            |
| `PlmMetadataFormatDate`           | DATE                              |
| `PlmMetadataFormatBoolean`        | BOOLEAN                           |
| `PlmMetadataFormatFile`           | FILE                              |
| `PlmMetadataFormatTextarea`       | TEXTAREA                          |
| `PlmMetadataFormatSingleSelect`   | SINGLE_SELECT                     |
| `PlmMetadataFormatMultipleSelect` | MULTIPLE_SELECT                   |
| `PlmMetadataFormatMeasurement`    | MEASUREMENT                       |
| `PlmMetadataFormatPrice`          | PRICE                             |
| `PlmMetadataFormatYesNo`          | YES_NO                            |
| `PlmMetadataFormatCurrency`       | CURRENCY                          |
| `PlmMetadataFormatURL`            | URL                               |
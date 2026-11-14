# QmsMetadataFormat

The data type of the custom field, which determines how `value` is formatted.

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.QmsMetadataFormatText

// Open enum: custom values can be created with a direct type cast
custom := shared.QmsMetadataFormat("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `QmsMetadataFormatText`           | TEXT                              |
| `QmsMetadataFormatNumber`         | NUMBER                            |
| `QmsMetadataFormatDate`           | DATE                              |
| `QmsMetadataFormatBoolean`        | BOOLEAN                           |
| `QmsMetadataFormatFile`           | FILE                              |
| `QmsMetadataFormatTextarea`       | TEXTAREA                          |
| `QmsMetadataFormatSingleSelect`   | SINGLE_SELECT                     |
| `QmsMetadataFormatMultipleSelect` | MULTIPLE_SELECT                   |
| `QmsMetadataFormatMeasurement`    | MEASUREMENT                       |
| `QmsMetadataFormatPrice`          | PRICE                             |
| `QmsMetadataFormatYesNo`          | YES_NO                            |
| `QmsMetadataFormatCurrency`       | CURRENCY                          |
| `QmsMetadataFormatURL`            | URL                               |
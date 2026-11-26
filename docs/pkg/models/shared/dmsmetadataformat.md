# DmsMetadataFormat

The data type of the custom field, which determines how `value` is formatted.

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.DmsMetadataFormatText

// Open enum: custom values can be created with a direct type cast
custom := shared.DmsMetadataFormat("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `DmsMetadataFormatText`           | TEXT                              |
| `DmsMetadataFormatNumber`         | NUMBER                            |
| `DmsMetadataFormatDate`           | DATE                              |
| `DmsMetadataFormatBoolean`        | BOOLEAN                           |
| `DmsMetadataFormatFile`           | FILE                              |
| `DmsMetadataFormatTextarea`       | TEXTAREA                          |
| `DmsMetadataFormatSingleSelect`   | SINGLE_SELECT                     |
| `DmsMetadataFormatMultipleSelect` | MULTIPLE_SELECT                   |
| `DmsMetadataFormatMeasurement`    | MEASUREMENT                       |
| `DmsMetadataFormatPrice`          | PRICE                             |
| `DmsMetadataFormatYesNo`          | YES_NO                            |
| `DmsMetadataFormatCurrency`       | CURRENCY                          |
| `DmsMetadataFormatURL`            | URL                               |
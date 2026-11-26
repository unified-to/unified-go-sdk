# PrintingMetadataFormat

The data type of the custom field, which determines how `value` is formatted.

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PrintingMetadataFormatText

// Open enum: custom values can be created with a direct type cast
custom := shared.PrintingMetadataFormat("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `PrintingMetadataFormatText`           | TEXT                                   |
| `PrintingMetadataFormatNumber`         | NUMBER                                 |
| `PrintingMetadataFormatDate`           | DATE                                   |
| `PrintingMetadataFormatBoolean`        | BOOLEAN                                |
| `PrintingMetadataFormatFile`           | FILE                                   |
| `PrintingMetadataFormatTextarea`       | TEXTAREA                               |
| `PrintingMetadataFormatSingleSelect`   | SINGLE_SELECT                          |
| `PrintingMetadataFormatMultipleSelect` | MULTIPLE_SELECT                        |
| `PrintingMetadataFormatMeasurement`    | MEASUREMENT                            |
| `PrintingMetadataFormatPrice`          | PRICE                                  |
| `PrintingMetadataFormatYesNo`          | YES_NO                                 |
| `PrintingMetadataFormatCurrency`       | CURRENCY                               |
| `PrintingMetadataFormatURL`            | URL                                    |
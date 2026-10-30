# PropertyMetadataFormat

the data type of the value

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.PropertyMetadataFormatText

// Open enum: custom values can be created with a direct type cast
custom := shared.PropertyMetadataFormat("custom_value")
```


## Values

| Name                                   | Value                                  |
| -------------------------------------- | -------------------------------------- |
| `PropertyMetadataFormatText`           | TEXT                                   |
| `PropertyMetadataFormatNumber`         | NUMBER                                 |
| `PropertyMetadataFormatDate`           | DATE                                   |
| `PropertyMetadataFormatBoolean`        | BOOLEAN                                |
| `PropertyMetadataFormatFile`           | FILE                                   |
| `PropertyMetadataFormatTextarea`       | TEXTAREA                               |
| `PropertyMetadataFormatSingleSelect`   | SINGLE_SELECT                          |
| `PropertyMetadataFormatMultipleSelect` | MULTIPLE_SELECT                        |
| `PropertyMetadataFormatMeasurement`    | MEASUREMENT                            |
| `PropertyMetadataFormatPrice`          | PRICE                                  |
| `PropertyMetadataFormatYesNo`          | YES_NO                                 |
| `PropertyMetadataFormatCurrency`       | CURRENCY                               |
| `PropertyMetadataFormatURL`            | URL                                    |
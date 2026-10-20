# SigningMetadataFormat

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.SigningMetadataFormatText

// Open enum: custom values can be created with a direct type cast
custom := shared.SigningMetadataFormat("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `SigningMetadataFormatText`           | TEXT                                  |
| `SigningMetadataFormatNumber`         | NUMBER                                |
| `SigningMetadataFormatDate`           | DATE                                  |
| `SigningMetadataFormatBoolean`        | BOOLEAN                               |
| `SigningMetadataFormatFile`           | FILE                                  |
| `SigningMetadataFormatTextarea`       | TEXTAREA                              |
| `SigningMetadataFormatSingleSelect`   | SINGLE_SELECT                         |
| `SigningMetadataFormatMultipleSelect` | MULTIPLE_SELECT                       |
| `SigningMetadataFormatMeasurement`    | MEASUREMENT                           |
| `SigningMetadataFormatPrice`          | PRICE                                 |
| `SigningMetadataFormatYesNo`          | YES_NO                                |
| `SigningMetadataFormatCurrency`       | CURRENCY                              |
| `SigningMetadataFormatURL`            | URL                                   |
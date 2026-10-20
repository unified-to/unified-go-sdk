# MessagingMetadataFormat

## Example Usage

```go
import (
	"github.com/unified-to/unified-go-sdk/pkg/models/shared"
)

value := shared.MessagingMetadataFormatText

// Open enum: custom values can be created with a direct type cast
custom := shared.MessagingMetadataFormat("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `MessagingMetadataFormatText`           | TEXT                                    |
| `MessagingMetadataFormatNumber`         | NUMBER                                  |
| `MessagingMetadataFormatDate`           | DATE                                    |
| `MessagingMetadataFormatBoolean`        | BOOLEAN                                 |
| `MessagingMetadataFormatFile`           | FILE                                    |
| `MessagingMetadataFormatTextarea`       | TEXTAREA                                |
| `MessagingMetadataFormatSingleSelect`   | SINGLE_SELECT                           |
| `MessagingMetadataFormatMultipleSelect` | MULTIPLE_SELECT                         |
| `MessagingMetadataFormatMeasurement`    | MEASUREMENT                             |
| `MessagingMetadataFormatPrice`          | PRICE                                   |
| `MessagingMetadataFormatYesNo`          | YES_NO                                  |
| `MessagingMetadataFormatCurrency`       | CURRENCY                                |
| `MessagingMetadataFormatURL`            | URL                                     |